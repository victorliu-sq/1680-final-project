package command

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
)

// ************************************************************************
// Hello Msg struct
type HelloMsg struct {
	CommandType uint8
	UdpPort     uint16
}

func NewHelloMsg(portNumS string) *HelloMsg {
	portNum, _ := strconv.Atoi(portNumS)
	helloMsg := &HelloMsg{
		CommandType: TypeHello,
		UdpPort:     uint16(portNum),
	}
	return helloMsg
}

func (helloMsg *HelloMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, helloMsg.CommandType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, helloMsg.UdpPort)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func UnmarsharlHelloMsg(bytes []byte) *HelloMsg {
	// buf := new(bytes.Buffer)
	// err := binary.Write(buf, binary.BigEndian, helloMsg.CommandType)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// err = binary.Write(buf, binary.BigEndian, helloMsg.UdpPort)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// return buf.Bytes()
	commandType := uint8(bytes[0])
	udpPort := uint16(binary.BigEndian.Uint16(bytes[1:]))
	helloMsg := &HelloMsg{
		CommandType: commandType,
		UdpPort:     udpPort,
	}
	return helloMsg
}
