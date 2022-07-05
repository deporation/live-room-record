package main

import (
	"live-room/config"
	"live-room/http"
	"live-room/service"
)

func main() {

	conf := config.GetInstance()

	clientMap := make(map[int]*service.BliveClient)

	for _, val := range conf.LiveRoom {
		go http.ListenRoomStart(val, 10)
	}

	for {
		select {
		case roomId := <-http.StartChannel:
			client := &service.BliveClient{}
			clientMap[roomId] = client
			handler := service.Init()
			go handler.HandlerInsert()
			go clientMap[roomId].Start(roomId, handler)
		case roomId := <-http.StopChannel:
			if clientMap[roomId] != nil {
				clientMap[roomId].Close()
				delete(clientMap, roomId)
			}
		}
	}
}
