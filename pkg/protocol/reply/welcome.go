package reply

import (
	"bytes"
	"encoding/binary"
	"log"
)

type WelcomeMsg struct {
	replyType   uint8
	numStations uint16
}

func NewWelcomeMsg(numStations uint16) *WelcomeMsg {
	wMsg := &WelcomeMsg{
		replyType:   TypeWelcome,
		numStations: numStations,
	}
	return wMsg
}

func (wMsg *WelcomeMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, wMsg.replyType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, wMsg.numStations)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func UnmarsharlWelcomeMsg(bytes []byte) *WelcomeMsg {
	wMsg := &WelcomeMsg{
		replyType:   uint8(bytes[0]),
		numStations: uint16(binary.BigEndian.Uint16(bytes[1:])),
	}
	return wMsg
}

func (wMsg *WelcomeMsg) GetNumStations() uint16 {
	return wMsg.numStations
}
