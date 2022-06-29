package model

import "context"

type ReceiveMessage struct {
	WsHeader

	Body []byte
}

type Context struct {
	context.Context
	Operation uint32
	Buffer    []byte
}

func NewMessage(data []byte) *ReceiveMessage {
	var msg ReceiveMessage
	msg.Body = data
	return &msg
}
