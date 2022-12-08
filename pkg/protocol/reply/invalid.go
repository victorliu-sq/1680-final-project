package reply

import (
	"bytes"
	"encoding/binary"
	"log"
)

type InvalidCommandMsg struct {
	ReplyType       uint8
	ReplyStringSize uint8
	ReplyString     []byte
}

func NewInvalidCommandMsg(replyStringS string) *InvalidCommandMsg {
	n := len(replyStringS)
	replyString := []byte(replyStringS)
	announceMsg := &InvalidCommandMsg{
		ReplyType:       TypeInvalid,
		ReplyStringSize: uint8(n),
		ReplyString:     replyString,
	}
	return announceMsg
}

func (invalidCommandMsg *InvalidCommandMsg) Marsharl() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, invalidCommandMsg.ReplyType)
	if err != nil {
		log.Fatalln(err)
	}
	err = binary.Write(buf, binary.BigEndian, invalidCommandMsg.ReplyStringSize)
	if err != nil {
		log.Fatalln(err)
	}
	bytes := buf.Bytes()
	bytes = append(bytes, invalidCommandMsg.ReplyString...)
	return bytes
}
