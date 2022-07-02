package service

import (
	"testing"
	"time"
)

func TestMengci(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(30 * time.Minute)
		client.Close()
	}()
	err := client.Start(2267110, &BaseHandler{})
	if err != nil {
		return
	}
}
func TestHair(t *testing.T) {
	client := BliveClient{}
	go func() {
		time.Sleep(30 * time.Minute)
		client.Close()
	}()
	err := client.Start(14343955, &BaseHandler{})
	if err != nil {
		return
	}
}
func TestHuahua(t *testing.T) {

	client := BliveClient{}
	go func() {
		time.Sleep(30 * time.Minute)
		client.Close()
	}()
	err := client.Start(7688602, &BaseHandler{})
	if err != nil {
		return
	}
}
