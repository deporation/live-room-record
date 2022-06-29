package websocket

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"io"
	"live-room/model"
	"live-room/util"
	"log"
	"testing"
	"time"
)

func TestWsClient(t *testing.T) {

	cer := model.Certificate{Uid: 0, RoomId: model.RoomId, Protover: 3, Platform: "web", Type: 2}
	marshal, err := json.Marshal(cer)
	if err != nil {
		t.Error(err)
		return
	}

	buffer, err := util.SetHeader(model.AUTH)
	if err != nil {
		return
	}
	ws, _, err := websocket.DefaultDialer.Dial("wss://ks-live-dmcmt-sh2-pm-03.chat.bilibili.com:443/sub", nil)

	err = ws.WriteMessage(websocket.BinaryMessage, append(buffer.Bytes(), marshal...))
	if err != nil {
		ws.Close()
		t.Error(err)
		log.Println(err)
		return
	}
	go func() {
		for {
			time.Sleep(time.Second * 30)
			buffer, err := util.SetHeader(model.HEARTBEAT)
			err = ws.WriteMessage(websocket.BinaryMessage, append(buffer.Bytes(), nil...))
			if err != nil {
				ws.Close()
				t.Error(err)
				log.Println(err)
				return
			}
		}
	}()
	for {
		// 等待信息返回
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			t.Error(err)
			ws.Close()
			return
		}
		if messageType != websocket.BinaryMessage {
			log.Printf("room=%d unknown websocket message type=%v, data=%s", model.RoomId, messageType, string(message))
		}

		log.Println(message[0:16], fmt.Sprintf("%s", string(message[16:])))
		//0 1 2 3 4 5 6 7 8
		//header, err := util.ReadHeader(message)
		if err != nil {
			return
		}
		headByte := bytes.NewBuffer(message[0:4])
		var headint32 int32
		err = binary.Read(headByte, binary.BigEndian, &headint32)
		if err != nil {
			t.Error(err)
			ws.Close()
			return
		}
		bytesBuffer := bytes.NewBuffer(message[8:12])
		var ageen int32
		err = binary.Read(bytesBuffer, binary.BigEndian, &ageen)
		if err != nil {
			log.Println(err, 222)
			return
		}
		switch ageen {
		case 2:
			b := bytes.NewReader(message[16:])
			r, _ := zlib.NewReader(b)
			bs, _ := io.ReadAll(r)
			log.Printf("zip压缩: %s", string(bs))
		case 0:
			b := message[16:]
			log.Println("json弹幕", string(b))
		case 3:
			b := message[16:]
			var renqiInt int32
			bytesBuffers := bytes.NewBuffer(b)
			err = binary.Read(bytesBuffers, binary.BigEndian, &renqiInt)
			if err != nil {
				log.Println(err, 222)
				return
			}
			log.Println("人气：", renqiInt)
		case 10:
			b := bytes.NewReader(message[16:])
			r := brotli.NewReader(b)
			bytess, _ := io.ReadAll(r)
			log.Println("brotli压缩", string(bytess), "解压前长度:", len(message[16:]), "解压长度:", len(bytess))
		default:
			log.Println("其他:", ageen)
		}

	}
}
