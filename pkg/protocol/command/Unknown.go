package command

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
)

// ************************************************************************
// Hello Msg struct
type UnknownMsg struct {
	CommandType uint8
	UdpPort     uint16
}

func NewUnknownMsg(portNumS string) *UnknownMsg {
	portNum, _ := strconv.Atoi(portNumS)
	unknownMsg := &UnknownMsg{
		CommandType: TypeUnknown,
		UdpPort:     uint16(portNum),
	}
	return unknownMsg
}

func (unknownMsg *UnknownMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, unknownMsg.CommandType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, unknownMsg.UdpPort)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}
