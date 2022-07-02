package service

import (
	"fmt"
	"live-room/model"
	"runtime"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	num := runtime.NumCPU() //这里得到当前所有的cpu数量
	fmt.Printf("couNum=%v\n", num)
	client := BliveClient{}
	go func() {
		time.Sleep(30 * time.Minute)
		client.Close()
	}()
	err := client.Start(model.RoomId, &BaseHandler{})
	if err != nil {
		return
	}
}
func TestLovely(t *testing.T) {
	num := runtime.NumCPU() //这里得到当前所有的cpu数量
	fmt.Printf("couNum=%v\n", num)
	client := BliveClient{}
	go func() {
		time.Sleep(30 * time.Minute)
		client.Close()
	}()
	err := client.Start(21692711, &BaseHandler{})
	if err != nil {
		return
	}
}
