package util

import (
	"bytes"
	"encoding/binary"
	"github.com/deporation/live-room-record/model"
)

func SetHeader(operate uint32, data []byte) (*bytes.Buffer, error) {
	header := model.WsHeader{
		PacketLen: uint32(binary.Size(data) + 16),
		HeaderLen: 16,
		Version:   1,
		Operation: operate,
		Sequence:  1,
	}
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, header)
	return buffer, nil
}

// ReadHeader 读取头
func ReadHeader(data []byte) (*model.WsHeader, error) {
	buffer := bytes.NewBuffer(data)
	header := model.WsHeader{}
	err := binary.Read(buffer, binary.BigEndian, &header)
	return &header, err
}
