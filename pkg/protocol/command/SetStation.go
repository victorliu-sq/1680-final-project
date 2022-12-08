package command

import (
	"bytes"
	"encoding/binary"
	"log"
)

// ************************************************************************
// Hello Msg struct
type SetStationMsg struct {
	CommandType   uint8
	StationNumber uint16
}

func NewSetStationMsg(stationNum uint16) *SetStationMsg {
	setStationMsg := &SetStationMsg{
		CommandType:   TypeSetStation,
		StationNumber: stationNum,
	}
	return setStationMsg
}

func (setStation *SetStationMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, setStation.CommandType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, setStation.StationNumber)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func UnmarshalSetstationMsg(bytes []byte) *SetStationMsg {
	setStationMsg := &SetStationMsg{
		CommandType:   uint8(bytes[0]),
		StationNumber: uint16(binary.BigEndian.Uint16(bytes[1:])),
	}
	return setStationMsg
}
