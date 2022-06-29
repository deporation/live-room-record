package http

import (
	"fmt"
	httpPro "net/http"
	"testing"
)

type biRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Group            string  `json:"group"`
		BusinessId       int     `json:"business_id"`
		RefreshRowFactor float64 `json:"refresh_row_factor"`
		RefreshRate      int     `json:"refresh_rate"`
		MaxDelay         int     `json:"max_delay"`
		Token            string  `json:"token"`
		HostList         []struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			WssPort int    `json:"wss_port"`
			WsPort  int    `json:"ws_port"`
		} `json:"host_list"`
	} `json:"data"`
}

func TestHttpRoomInit(t *testing.T) {
	header := httpPro.Header{}
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	var test biRes
	err := Get("https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo", nil, &header, map[string]string{"id": "7688602", "type": "0"}, &test)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(test)

}
