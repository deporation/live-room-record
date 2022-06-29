package model

type WsHeader struct {
	Length        *int32
	RawHeaderSize *int16
	Ver           *int16
	Operate       *int32
	Seq           *int32
}
