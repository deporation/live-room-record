package service

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/andybalholm/brotli"
	httpClient "github.com/bilibili/live-room-record/http"
	"github.com/bilibili/live-room-record/model"
	"github.com/bilibili/live-room-record/util"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"strconv"
	"time"
)

const (
	ROOM_INIT_URL           = "https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom"
	DANMAKU_SERVER_CONF_URL = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
)

const (
	DANMU_MSG                     = "DANMU_MSG"
	SEND_GIFT                     = "SEND_GIFT"
	SUPER_CHAT_MESSAGE            = "SUPER_CHAT_MESSAGE"
	GUARD_BUY                     = "GUARD_BUY"
	ROOM_RANK                     = "ROOM_RANK"
	ROOM_REAL_TIME_MESSAGE_UPDATE = "ROOM_REAL_TIME_MESSAGE_UPDATE"
	ROOM_CHANGE                   = "ROOM_CHANGE"
	INTERACT_WORD                 = "INTERACT_WORD"
	ONLINE_RANK_V2                = "ONLINE_RANK_V2"
	ONLINE_RANK_TOP3              = "ONLINE_RANK_TOP3"
	ONLINE_RANK_COUNT             = "ONLINE_RANK_COUNT"
	COMBO_SEND                    = "COMBO_SEND"
	WIDGET_BANNER                 = "WIDGET_BANNER"
	ENTRY_EFFECT                  = "ENTRY_EFFECT"
	RQZ                           = "RQZ"
	DEFAULT                       = "DEFAULT"
	LIVE                          = "LIVE"
	PREPARING                     = "PREPARING"
)

type BliveClient struct {
	roomId            int                        // URL中的房间ID
	shortId           int                        //直播间短id
	uid               int                        //uid: B站用户ID，0表示未登录
	heartbeatInterval int                        //发送心跳包的间隔时间（秒）
	handler           Handler                    // 事件处理器
	roomerOwnerUid    int                        //主播用户ID
	conn              *websocket.Conn            //websocket连接
	rawQueue          chan *model.ReceiveMessage //ws 接收数据
	msgQueue          chan *model.Context        // 处理队列
	close             chan struct{}              // close client connection
	heartBeatErr      chan error                 //心跳发送error
	recieveErr        chan error                 //接收message的error
	ctx               context.Context
	Status            int //客户端状况
}

func (client *BliveClient) clientInit() {
	//head := &http.Header{}
	var roomInfo model.RoomInfo
	err := httpClient.Get(ROOM_INIT_URL, nil, nil, map[string]string{"room_id": strconv.Itoa(client.roomId)}, &roomInfo)
	if err != nil && &roomInfo == nil {
		panic(err)
		return
	}
	client.roomId = roomInfo.Data.RoomInfo.RoomId
	client.shortId = roomInfo.Data.RoomInfo.ShortId
	client.roomerOwnerUid = roomInfo.Data.RoomInfo.Uid
	client.heartbeatInterval = 30
	client.close = make(chan struct{}, 1)
	client.rawQueue = make(chan *model.ReceiveMessage, 10)
	client.msgQueue = make(chan *model.Context, 10)
	client.heartBeatErr = make(chan error, 1)
	client.recieveErr = make(chan error, 1)
	var danmuServer model.BiRes
	err = httpClient.Get(DANMAKU_SERVER_CONF_URL, nil, nil, map[string]string{"id": strconv.Itoa(client.roomId), "type": "0"}, &danmuServer)
	if err != nil {
		panic(err)
		return
	}
	for index, url := range danmuServer.Data.HostList {
		client.conn, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("wss://%s:%d/sub", "broadcastlv.chat.bilibili.com", url.WssPort), nil)
		if err != nil {
			log.Printf("%dhost:%s，无法连接", index, url.Host)
			continue
		} else {
			break
		}
	}
	auth := model.Auth{
		UID:      0,
		Roomid:   uint32(client.roomId),
		Protover: 2,
		Platform: "web",
		Type:     2,
		Key:      danmuServer.Data.Token,
	}
	marshal, err := json.Marshal(auth)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.send(marshal, model.AUTH)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (client *BliveClient) Close() {
	close(client.heartBeatErr)
	close(client.recieveErr)
	close(client.rawQueue)
	close(client.msgQueue)
	client.conn.Close()
	client.close <- struct{}{}
}

func (client *BliveClient) Restart() {
	go client.Start(client.roomId, client.handler)
}

func (client *BliveClient) Start(roomId int, handler Handler) error {
	client.roomId = roomId
	client.clientInit()
	client.handler = handler
	ctx, cancel := context.WithCancel(context.Background())
	client.ctx = ctx
	defer cancel()

	go client.heartBeat(ctx)
	go client.recieve(ctx)
	go client.parse(ctx)
	go client.handle(ctx)

	client.Status = 1

	select {
	case err := <-client.heartBeatErr:
		client.Status = -1
		return err
	case err := <-client.recieveErr:
		client.Status = -1
		return err
	case <-ctx.Done():
	case <-client.close:
		return nil
	}
	return nil
}

// websocket send message
// data: []byte 传输数据
// operate: uint32 操作码
func (client *BliveClient) send(data []byte, operate uint32) error {
	message, err := util.SetHeader(operate, data)

	if err != nil {
		return err
	}

	sendData := append(message.Bytes(), data...)

	return client.conn.WriteMessage(websocket.BinaryMessage, sendData)

}

// 定时发送心跳
func (client *BliveClient) heartBeat(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-client.close:
			return
		default:
			if err := client.send([]byte(""), model.HEARTBEAT); err != nil {
				log.Println("heart error", err)
				if client.Status == 0 {
					return
				}

				client.heartBeatErr <- err
				return
			}
		}
		time.Sleep(time.Second * time.Duration(client.heartbeatInterval))
	}
}

func (client *BliveClient) recieve(ctx context.Context) {
	count := 0
	closeConn := make(chan struct{}, 1)
	for {
		select {
		case <-ctx.Done():
		case <-client.close:
			return
		case <-closeConn:
			log.Println("reconnecting...")
			client.clientInit()
			count = 0
		default:
			_, body, err := client.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Unexpected Close Error")
					closeConn <- struct{}{}
					continue
				}
				if count != 5 {
					count++
					continue
				}
				log.Printf("RoomID=%d Read Message Error...", client.roomId)
				client.recieveErr <- err
				return
			}
			msg := &model.ReceiveMessage{Body: body}
			if client.Status == 0 {
				return
			}
			client.rawQueue <- msg
		}
	}
}

func (client *BliveClient) parse(ctx context.Context) {
	var (
		msg          *model.ReceiveMessage
		header       model.WsHeader
		headerBuffer *bytes.Reader
		buffer       []byte
	)
	for {
		msg = <-client.rawQueue
		for msg != nil && len(msg.Body) > 0 {
			select {
			case <-ctx.Done():
			case <-client.close:
				return
			default:
			}

			headerBuffer = bytes.NewReader(msg.Body[:16])
			_ = binary.Read(headerBuffer, binary.BigEndian, &header)
			buffer = msg.Body[16:int(header.PacketLen)]
			msg.Body = msg.Body[int(header.PacketLen):]

			if int(header.Version) == DEFLATE {
				msg.Body = zlibInflate(buffer)
				continue
			}

			if header.Version == BROTLI {
				msg.Body = brotliInflate(buffer)
				continue
			}
			if len(buffer) > 0 {
				if client.Status == 0 {
					return
				}
				client.msgQueue <- &model.Context{Context: ctx, Operation: header.Operation, Buffer: buffer}
			}
		}

	}
}

func (client *BliveClient) handle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		case <-client.close:
			return
		default:
		}
		msg := <-client.msgQueue
		if msg == nil {
			return
		}
		buffer := bytes.NewReader(msg.Buffer)

		switch msg.Operation {
		case model.HEARTBEAT_REPLY:
			var popular int32
			_ = binary.Read(buffer, binary.BigEndian, &popular)
			go client.handler.hearBeat(ctx, model.HeartbeatMessage{Popularity: popular, Roomid: client.roomId})
			break
		case model.SEND_MSG_REPLY:
			message := gjson.GetBytes(msg.Buffer, "cmd").String()
			switch message {
			case DANMU_MSG:
				info := gjson.GetBytes(msg.Buffer, "info").Array()
				danmuMessage := model.ParaseMessage(info)
				danmuMessage.RoomId = client.roomId
				go client.handler.danmuku(ctx, *danmuMessage)
				break
			case SEND_GIFT:
				var gift model.GiftMessage
				err := json.Unmarshal(msg.Buffer, &gift)
				if err != nil {
					log.Printf("json 异常")
					return
				}
				gift.RoomId = client.roomId
				go client.handler.sendGift(ctx, gift)
				break
			case SUPER_CHAT_MESSAGE:
				var sc model.SuperChatMessage
				err := json.Unmarshal(msg.Buffer, &sc)
				if err != nil {
					log.Printf("json 异常")
					return
				}
				sc.Roomid = client.roomId
				go client.handler.superChat(ctx, sc)
				break
			case GUARD_BUY:
				log.Println(string(msg.Buffer))
				var guard model.GuardMessage
				err := json.Unmarshal(msg.Buffer, &guard)
				if err != nil {
					log.Printf("json 异常")
					return
				}
				guard.RoomId = client.roomId
				go client.handler.buyGuard(ctx, guard)
				break
			case LIVE:
				log.Println(LIVE, string(msg.Buffer))
				break
			case ROOM_REAL_TIME_MESSAGE_UPDATE:
				break
			case ROOM_CHANGE:
				break
			}

		default:
			log.Println(string(msg.Buffer))
			break
		}
	}
}

func zlibInflate(src []byte) []byte {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func brotliInflate(src []byte) []byte {
	b := bytes.NewReader(src)
	r := brotli.NewReader(b)
	res, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return res
}
