package reply

import (
	"bytes"
	"encoding/binary"
	"log"
)

type StationShutDown struct {
	ReplyType  uint8
	StationIdx uint16
}

func NewStationShutDown(stationIdx uint16) *StationShutDown {
	sMsg := &StationShutDown{
		ReplyType:  TypeStationShutDown,
		StationIdx: stationIdx,
	}
	return sMsg
}

func (sMsg *StationShutDown) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, sMsg.ReplyType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, sMsg.StationIdx)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func UnmarsharlStationShutDownMsg(bytes []byte) *StationShutDown {
	sMsg := &StationShutDown{
		ReplyType:  uint8(bytes[0]),
		StationIdx: uint16(binary.BigEndian.Uint16(bytes[1:])),
	}
	return sMsg
}
