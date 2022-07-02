package main

import (
	"live-room/model"
	"live-room/service"
)

func main() {
	client := service.BliveClient{}
	client.Start(model.RoomId, &service.BaseHandler{})
	defer client.Close()
}
