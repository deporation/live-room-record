package service

import (
	"live-room/model"
	"log"
)

type Handler interface {
	hearBeat(message model.HeartbeatMessage) //心跳
	danmuku()                                //弹幕库
	buyGuard()                               // 航海服务
	superChat()                              // sc
}

type BaseHandler struct{}

func (handler *BaseHandler) hearBeat(message model.HeartbeatMessage) {
	log.Printf("人气：%d", message.Popularity)
}

func (handler *BaseHandler) danmuku() {

}

func (handler *BaseHandler) buyGuard() {

}

func (handler *BaseHandler) superChat() {

}
