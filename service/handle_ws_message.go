package service

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"live-room/model"
	"log"
	"strconv"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
)

type Handler interface {
	hearBeat(ctx context.Context, message model.HeartbeatMessage) //心跳
	danmuku(ctx context.Context, message model.Message)           //弹幕库
	sendGift(ctx context.Context, message model.GiftMessage)
	buyGuard(ctx context.Context, message model.GuardMessage)      // 航海服务
	superChat(ctx context.Context, message model.SuperChatMessage) // sc
}

type BaseHandler struct{}

func (handler *BaseHandler) hearBeat(ctx context.Context, message model.HeartbeatMessage) {
	log.Printf("人气：%d", message.Popularity)
}

func (handler *BaseHandler) danmuku(ctx context.Context, message model.Message) {
	log.Printf("%s, %s ,%s: %s", green(message.Uid), blue(message.MedalName), yellow(message.Uname), message.Msg)
}

func (handler *BaseHandler) sendGift(ctx context.Context, message model.GiftMessage) {
	log.Printf("%s, %s ,:%s, %s", green(message.Data.UID), blue(message.Data.Uname), yellow(message.Data.GiftName+" * "+strconv.Itoa(message.Data.Num)), message.Data.CoinType+" * "+strconv.Itoa(message.Data.TotalCoin))
}

func (handler *BaseHandler) buyGuard(ctx context.Context, message model.GuardMessage) {
	log.Printf("%s, %s ,:%s, %d 天", green(message.Data.Uid), blue(message.Data.UserName), yellow(message.Data.GiftName+" * "+strconv.Itoa(message.Data.Num)), (message.Data.EndTime-message.Data.StartTime)/(60*60*24/1000))
}

func (handler *BaseHandler) superChat(ctx context.Context, message model.SuperChatMessage) {
	log.Printf("roomId: %s , %s, %s ,:%s", strconv.Itoa(message.Roomid), green(message.Data.Uid), blue(message.Data.UserInfo.Uname), yellow(fmt.Sprintf("醒目留言:%s", message.Data.Message)))
}
