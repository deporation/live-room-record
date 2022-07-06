package main

import (
	"github.com/bilibili/live-room-record/config"
	"github.com/bilibili/live-room-record/http"
	"github.com/bilibili/live-room-record/service"
	"log"
	"sync"
)

func main() {

	conf := config.GetInstance()

	clientMap := make(map[int]*service.BliveClient)

	for _, val := range conf.LiveRoom {
		go http.ListenRoomStart(val, 10)
	}

	var lockStop sync.Mutex
	for {
		select {
		case roomId := <-http.StartChannel:
			lockStop.Lock()
			client := service.BliveClient{}
			clientMap[roomId] = &client
			handler := service.Init()
			go handler.HandlerInsert()
			go func() {
				err := clientMap[roomId].Start(roomId, handler)
				if err != nil {
					return
				}
			}()
			lockStop.Unlock()
		case roomId := <-http.StopChannel:
			lockStop.Lock()
			if v, ok := clientMap[roomId]; ok {
				v.Status = 0
				log.Println("开始close", roomId)
				v.Close()
				log.Println("close  结束", roomId)
				delete(clientMap, roomId)
			}
			lockStop.Unlock()
		default:
			for _, value := range clientMap {
				if value.Status == -1 {
					value.Restart()
				}
			}
		}
	}
}
