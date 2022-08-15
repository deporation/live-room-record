# 记录bilibili直播间信息
可根据配置直播间号来记录直播间信息  
默认采用MongoDB来保存信息  
可以通过实现[Handler](./service/handle_ws_message.go)接口,
来进行自定义  
```go
type Handler interface {
	heartBeat(ctx context.Context, message model.HeartbeatMessage)  //心跳
	danmuku(ctx context.Context, message model.Message)            //弹幕库
	sendGift(ctx context.Context, message model.GiftMessage)       // 礼物
	buyGuard(ctx context.Context, message model.GuardMessage)      // 航海服务
	superChat(ctx context.Context, message model.SuperChatMessage) // sc
	Close(ctx context.Context)
}
```
具体参考[BaseHandler](./service/handle_ws_message.go)

仅做学习目的，商业行为概不负责