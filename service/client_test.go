package service

import (
	"testing"
	"time"
)

func TestMengci(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(6 * time.Hour)
		client.Close()
	}()
	handler := Init()
	go handler.HandlerInsert()
	err := client.Start(2267110, handler)
	if err != nil {
		return
	}
}
func TestXuehu(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(6 * time.Hour)
		client.Close()
	}()
	handler := Init()
	go handler.HandlerInsert()
	err := client.Start(24393, handler)
	if err != nil {
		return
	}
}
func TestXiyue(t *testing.T) {

	client := BliveClient{}
	go func() {
		time.Sleep(12 * time.Hour)
		client.Close()
	}()
	handler := Init()
	go handler.HandlerInsert()
	err := client.Start(22889484, handler)
	if err != nil {
		return
	}
}

func TestAiErSha(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(3 * time.Hour)
		client.Close()
	}()
	handler := Init()
	go handler.HandlerInsert()
	err := client.Start(81004, handler)
	if err != nil {
		return
	}
}

func TestTwoUncle(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(6 * time.Hour)
		client.Close()
	}()
	handler := Init()
	go handler.HandlerInsert()
	err := client.Start(12576972, handler)
	if err != nil {
		return
	}
}

func TestHuaHua(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(6 * time.Hour)
		client.Close()
	}()
	handler := Init()
	go handler.HandlerInsert()
	err := client.Start(7688602, handler)
	if err != nil {
		return
	}
}
