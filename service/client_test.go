package service

import (
	"fmt"
	"live-room/model"
	"runtime"
	"testing"
)

func TestClient(t *testing.T) {
	num := runtime.NumCPU() //这里得到当前所有的cpu数量
	fmt.Printf("couNum=%v\n", num)
	client := BliveClient{}
	client.Start(model.RoomId, &BaseHandler{})
}
