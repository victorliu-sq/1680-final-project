package reply

import (
	"bytes"
	"encoding/binary"
	"log"
)

type NewStationMsg struct {
	ReplyType  uint8
	StationNum uint16
}

func NewNewStationMsg(numStations uint16) *NewStationMsg {
	nMsg := &NewStationMsg{
		ReplyType:  TypeNewStation,
		StationNum: numStations,
	}
	return nMsg
}

func (nMsg *NewStationMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, nMsg.ReplyType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, nMsg.StationNum)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func UnmarsharlNewStationMsg(bytes []byte) *NewStationMsg {
	nMsg := &NewStationMsg{
		ReplyType:  uint8(bytes[0]),
		StationNum: uint16(binary.BigEndian.Uint16(bytes[1:])),
	}
	return nMsg
}
