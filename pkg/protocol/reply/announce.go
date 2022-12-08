package reply

import (
	"bytes"
	"encoding/binary"
	"log"
)

type AnnounceMsg struct {
	ReplyType    uint8
	SongnameSize uint8
	Songname     []byte
}

func NewAnnounceMsg(songnameS string) *AnnounceMsg {
	n := len(songnameS)
	songname := []byte(songnameS)
	announceMsg := &AnnounceMsg{
		ReplyType:    TypeAnnounce,
		SongnameSize: uint8(n),
		Songname:     songname,
	}
	return announceMsg
}

func (announceMsg *AnnounceMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, announceMsg.ReplyType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, announceMsg.SongnameSize)
	if err != nil {
		log.Fatalln(err)
	}
	bytes := buf.Bytes()
	bytes = append(bytes, announceMsg.Songname...)
	return bytes
}
