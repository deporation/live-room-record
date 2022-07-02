package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"live-room/model"
	"live-room/util"
	"log"
	"net/url"
	"time"
)

// ClientManager ws客户端
type ClientManager struct {
	conn        *websocket.Conn
	addr        *string
	path        string
	RecvMsgChan chan []byte
	isAlive     bool
	timeout     int
	token       string
}

func NewWsClientManager(addrIp, addrPort, path string, timeout int, token string) *ClientManager {
	addrString := addrIp + ":" + addrPort
	var recvChan = make(chan []byte, 100)
	var conn *websocket.Conn
	return &ClientManager{
		addr:        &addrString,
		path:        path,
		conn:        conn,
		RecvMsgChan: recvChan,
		isAlive:     false,
		timeout:     timeout,
	}
}

// 链接服务端
func (wsc *ClientManager) dail() {
	var err error
	u := url.URL{Scheme: "wss", Host: *wsc.addr, Path: wsc.path}
	log.Printf("connecting to %s", u.String())
	wsc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		fmt.Println(err)
		return

	}
	wsc.isAlive = true

	log.Printf("connecting to %s 链接成功！！！", u.String())
	cer := model.Certificate{Uid: 0, RoomId: model.RoomId, Protover: 3, Platform: "web", Type: 2, Key: wsc.token}
	marshal, err := json.Marshal(cer)
	if err != nil {
		log.Fatal(err)
		return
	}
	auth := model.Auth{
		UID:      0,
		Roomid:   model.RoomId,
		Protover: 2,
		Platform: "web",
		Type:     2,
		Key:      wsc.token,
	}
	marshal, err = json.Marshal(auth)
	if err != nil {
		log.Fatal(err)
		return
	}
	buffer, err := util.SetHeader(model.AUTH, marshal)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = wsc.conn.WriteMessage(websocket.BinaryMessage, append(buffer.Bytes(), marshal...))

}

// 发送消息
func (wsc *ClientManager) sendMsgThread() {
	go func() {
		for {
			buffer, err := util.SetHeader(model.HEARTBEAT, nil)
			err = wsc.conn.WriteMessage(websocket.BinaryMessage, append(buffer.Bytes(), nil...))
			if err != nil {
				log.Fatal(err)
				return
			}
			time.Sleep(time.Second * 30)
		}
	}()
}

// 读取消息
func (wsc *ClientManager) readMsgThread() {
	go func() {
		for {
			if wsc.conn != nil {

				binaryType, message, err := wsc.conn.ReadMessage()

				if err != nil {
					log.Println("read:", err)
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				if binaryType != websocket.BinaryMessage {
					log.Println("read:", errors.New("receive error"))
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				log.Printf("recv message: %s", fmt.Sprintf("%08x", message))
				// 需要读取数据，不然会阻塞
				wsc.RecvMsgChan <- message
			}

		}
	}()
}

// Start 开启服务并重连
func (wsc *ClientManager) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		default:
			if wsc.isAlive == false {
				wsc.dail()
				wsc.sendMsgThread()
				wsc.readMsgThread()
			}
		}
	}
}
