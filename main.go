package main

import (
	"github.com/deporation/live-room-record/config"
	"github.com/deporation/live-room-record/service"
)

func main() {

	conf := config.GetInstance()

	clientMap := make(map[int]*service.BliveClient)

	for _, roomId := range conf.LiveRoom {
		client := service.BliveClient{}
		clientMap[roomId] = &client
		handler := service.Init()
		go handler.HandlerInsert()
		go func(roomId int) {
			err := clientMap[roomId].Start(roomId, handler)
			if err != nil {
				return
			}
		}(roomId)
	}

	for {
		select {
		default:
			for _, value := range clientMap {
				if value.Status == -1 {
					value.Restart()
				}
			}
		}
	}
}
