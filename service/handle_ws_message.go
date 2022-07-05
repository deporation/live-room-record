package service

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"live-room/config"
	"live-room/library/orm"
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
	hearBeat(ctx context.Context, message model.HeartbeatMessage)  //心跳
	danmuku(ctx context.Context, message model.Message)            //弹幕库
	sendGift(ctx context.Context, message model.GiftMessage)       // 礼物
	buyGuard(ctx context.Context, message model.GuardMessage)      // 航海服务
	superChat(ctx context.Context, message model.SuperChatMessage) // sc
	Close(ctx context.Context)
}

type BaseHandler struct {
	mongodbs map[string]*orm.MongoClient

	danmuChan chan *model.Message
	giftChan  chan *model.GiftMessage
	scChan    chan *model.SuperChatMessage
	guardChan chan *model.GuardMessage
}

func Init() *BaseHandler {
	conf := config.GetInstance()
	danmuCollection := orm.MongoClient{}
	danmuCollection.Connect(conf.Momngo.Host, conf.Momngo.Port, conf.Momngo.Database, "danmu_ku")
	giftCollection := orm.MongoClient{}
	giftCollection.Connect(conf.Momngo.Host, conf.Momngo.Port, conf.Momngo.Database, "gift_ku")
	guardCollection := orm.MongoClient{}
	guardCollection.Connect(conf.Momngo.Host, conf.Momngo.Port, conf.Momngo.Database, "guard_ku")
	superChatCollection := orm.MongoClient{}
	superChatCollection.Connect(conf.Momngo.Host, conf.Momngo.Port, conf.Momngo.Database, "super_chat_ku")
	return &BaseHandler{
		mongodbs: map[string]*orm.MongoClient{
			"danmu": &danmuCollection,
			"gift":  &giftCollection,
			"guard": &guardCollection,
			"sc":    &superChatCollection,
		},
		danmuChan: make(chan *model.Message, 100),
		scChan:    make(chan *model.SuperChatMessage, 5),
		giftChan:  make(chan *model.GiftMessage, 20),
		guardChan: make(chan *model.GuardMessage, 5),
	}
}

func (handler *BaseHandler) hearBeat(ctx context.Context, message model.HeartbeatMessage) {
	log.Printf("人气：%d", message.Popularity)
}

func (handler *BaseHandler) danmuku(ctx context.Context, message model.Message) {
	handler.danmuChan <- &message
}

func (handler *BaseHandler) sendGift(ctx context.Context, message model.GiftMessage) {
	handler.giftChan <- &message
}

func (handler *BaseHandler) buyGuard(ctx context.Context, message model.GuardMessage) {

	handler.guardChan <- &message
}

func (handler *BaseHandler) superChat(ctx context.Context, message model.SuperChatMessage) {
	handler.scChan <- &message
}

func (handler *BaseHandler) Close(ctx context.Context) {
	close(handler.giftChan)
	close(handler.guardChan)
	close(handler.danmuChan)
	close(handler.scChan)
	for _, v := range handler.mongodbs {
		v.CLose(ctx)
	}
}

func (handler *BaseHandler) HandlerInsert() {

	for {

		select {
		case message := <-handler.giftChan:
			if message != nil {
				go func() {
					log.Printf("[%d] \t %s, %s ,:%s, %s", message.RoomId, green(message.Data.UID), blue(message.Data.Uname), yellow(message.Data.GiftName+" * "+strconv.Itoa(message.Data.Num)), message.Data.CoinType+" * "+strconv.Itoa(message.Data.TotalCoin))
					err, _ := handler.mongodbs["gift"].InsertedOne(context.TODO(), "", message)
					if err != nil {
						log.Println(err)
					}
				}()
			}
			break
		case message := <-handler.scChan:
			if message != nil {
				go func() {
					log.Printf("roomId: %s , %s, %s ,:%s", strconv.Itoa(message.Roomid), green(message.Data.Uid), blue(message.Data.UserInfo.Uname), yellow(fmt.Sprintf("醒目留言:%s", message.Data.Message)))
					err, _ := handler.mongodbs["sc"].InsertedOne(context.TODO(), "", message)
					if err != nil {
						log.Println(err)
					}
				}()
			}
			break
		case message := <-handler.danmuChan:
			if message != nil {
				go func() {
					log.Printf("[%d] \t %s, %s ,%s: %s", message.RoomId, green(message.Uid), blue(message.MedalName), yellow(message.Uname), message.Msg)
					err, _ := handler.mongodbs["danmu"].InsertedOne(context.TODO(), "", message)
					if err != nil {
						log.Println(err)
					}
				}()
			}
			break
		case message := <-handler.guardChan:
			if message != nil {
				go func() {
					log.Printf("[%d] %s, %s ,:%s, %d 天", message.RoomId, green(message.Data.Uid), blue(message.Data.UserName), yellow(message.Data.GiftName+" * "+strconv.Itoa(message.Data.Num)), (message.Data.EndTime-message.Data.StartTime)/(60*60*24/1000))
					err, _ := handler.mongodbs["guard"].InsertedOne(context.TODO(), "", message)
					if err != nil {
						log.Println(err)
					}
				}()
			}
			break
		default:
			break
		}
	}

}
