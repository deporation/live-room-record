package util

import (
	"bytes"
	"encoding/binary"
	"live-room/model"
	"log"
)

func SetHeader(operate int, data []byte) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer([]byte{})
	// 总长度 4字节
	if err := binary.Write(buffer, binary.BigEndian, int32(binary.Size(data)+16)); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 头部长度16 2字节
	if err := binary.Write(buffer, binary.BigEndian, int16(16)); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 协议版本号 固定1 2字节
	if err := binary.Write(buffer, binary.BigEndian, int16(1)); err != nil {
		log.Fatal(err)

		return nil, err
	}
	// 加入弹幕 固定协议7， 4字节
	if err := binary.Write(buffer, binary.BigEndian, int32(operate)); err != nil {
		log.Fatal(err)

		return nil, err
	}

	// 常量 1固定 ， 4字节
	if err := binary.Write(buffer, binary.BigEndian, int32(1)); err != nil {
		log.Fatal(err)

		return nil, err
	}
	return buffer, nil
}

// ReadHeader 读取头
func ReadHeader(data []byte) (*model.WsHeader, error) {
	buffer := bytes.NewBuffer(data)
	var length int32
	err := binary.Read(buffer, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}
	var rawHeaderSize int16
	err = binary.Read(buffer, binary.BigEndian, &rawHeaderSize)
	if err != nil {
		return nil, err
	}
	var ver int16
	err = binary.Read(buffer, binary.BigEndian, &ver)
	if err != nil {
		return nil, err
	}
	var operate int32
	err = binary.Read(buffer, binary.BigEndian, &operate)
	if err != nil {
		return nil, err
	}
	var seq int32
	err = binary.Read(buffer, binary.BigEndian, &seq)
	if err != nil {
		return nil, err
	}
	return &model.WsHeader{
		Length:        &length,
		RawHeaderSize: &rawHeaderSize,
		Ver:           &ver,
		Operate:       &operate,
		Seq:           &seq,
	}, nil
}
