package service

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/andybalholm/brotli"
	"io"
	httpClient "live-room/http"
	"live-room/model"
	"live-room/util"
	"live-room/websocket"
	"log"
	"strconv"
)

var i, j, k int

const (
	ROOM_INIT_URL           = "https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom"
	DANMAKU_SERVER_CONF_URL = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
)

type BliveClient struct {
	roomId            int                     // URL中的房间ID
	shortId           int                     //直播间短id
	uid               int                     //uid: B站用户ID，0表示未登录
	heartbeatInterval int                     //发送心跳包的间隔时间（秒）
	handler           Handler                 // 事件处理器
	roomerOwnerUid    int                     //主播用户ID
	webSocket         websocket.ClientManager //websocket连接
}

func (client *BliveClient) clientInit(roomId int64) {
	//head := &http.Header{}
	var roomInfo model.RoomInfo
	err := httpClient.Get(ROOM_INIT_URL, nil, nil, map[string]string{"room_id": strconv.FormatInt(roomId, 10)}, &roomInfo)
	if err != nil {
		panic(err)
		return
	}
	client.roomId = roomInfo.Data.RoomInfo.RoomId
	client.shortId = roomInfo.Data.RoomInfo.ShortId
	client.roomerOwnerUid = roomInfo.Data.RoomInfo.Uid

	var danmuServer model.BiRes
	err = httpClient.Get(DANMAKU_SERVER_CONF_URL, nil, nil, map[string]string{"id": strconv.FormatInt(roomId, 10), "type": "0"}, &danmuServer)
	if err != nil {
		panic(err)
		return
	}
	token := danmuServer.Data.Token
	port := danmuServer.Data.HostList[0].WssPort

	client.webSocket = *websocket.NewWsClientManager("broadcastlv.chat.bilibili.com", strconv.Itoa(port), "/sub", client.heartbeatInterval, token)
}

func (client *BliveClient) Start(roomId int64, handler Handler) {
	i = 0
	j = 0
	k = 0
	client.clientInit(roomId)
	client.handler = handler
	client.webSocket.Start()
	for {
		message, ok := <-client.webSocket.RecvMsgChan
		fmt.Println(ok, "is ok?")
		if ok {
			go client.handleMessage(message)
		} else {
			log.Println("ws no message")
		}
	}

}

func (client *BliveClient) handleMessage(message []byte) {
	j++
	fmt.Println(j)
	header, err := util.ReadHeader(message)
	if err != nil {
		return
	}

	offest := 0
	log.Printf("operate code is : %d", int(*header.Operate))
	switch int(*header.Operate) {
	case model.SEND_MSG_REPLY:
		for {
			data := message[offest+int(*header.RawHeaderSize) : offest+int(*header.Length)]

			go client.parseMessage(header, data)
			offest += int(*header.Length)
			log.Println("offest >= len", offest >= len(message))
			if offest >= len(message) {
				offest = 0
				break
			}
			header, err = util.ReadHeader(message[offest:])
			if err != nil {
				log.Fatal(err)
			}
		}
	case model.AUTH_REPLY:
		for {
			data := message[offest+int(*header.RawHeaderSize) : offest+int(*header.Length)]
			go client.parseMessage(header, data)
			offest = int(*header.Length) + offest
			if offest >= len(message) {
				offest = 0
				break
			}
			header, err = util.ReadHeader(message[offest:])
			if err != nil {
				log.Fatal(err)
			}
		}
		break
	case model.HEARTBEAT_REPLY:
		data := message[offest+int(*header.RawHeaderSize):]
		buffer := bytes.NewBuffer(data)
		var popular int32
		err = binary.Read(buffer, binary.BigEndian, &popular)
		go client.handler.hearBeat(model.HeartbeatMessage{Popularity: popular})
		break
	default:
		log.Println("not cache operate is :", int(*header.Operate))
		break

	}
	return
}

func (client *BliveClient) parseMessage(header *model.WsHeader, data []byte) {
	log.Printf("version:%d", *header.Ver)
	switch *header.Operate {
	case model.SEND_MSG_REPLY:
		if int(*header.Ver) == BROTLI {
			message := brotliInflate(data)
			client.webSocket.RecvMsgChan <- message
		} else if int(*header.Ver) == NORMAL {
			if len(data) != 0 {
				fmt.Println(string(data))
				log.Printf("data:%s", string(data))
			}
		} else if int(*header.Ver) == DEFLATE {
			message := zlibInflate(data)
			client.webSocket.RecvMsgChan <- message
		}
	default:
		return
	}
}

func zlibInflate(src []byte) []byte {
	i++
	fmt.Println("before zlib 解析", string(src))
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	fmt.Println("zlib 解析", string(out.Bytes()))
	defer r.Close()
	fmt.Println("zlib 调用:", i)
	return out.Bytes()
}

func brotliInflate(src []byte) []byte {
	k++
	b := bytes.NewReader(src)
	r := brotli.NewReader(b)
	res, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("brotli 解析", string(res))
	fmt.Println("brotli 调用:", k)
	return res
}
