package model

type WsHeader struct {
	PacketLen uint32 // 4
	HeaderLen uint16 // 2
	Version   uint16 // 2
	Operation uint32 // 4
	Sequence  uint32 // 4
}
