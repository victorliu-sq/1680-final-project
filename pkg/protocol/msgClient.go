package protocol

type ClientMsg struct {
	MsgType    uint8
	StationNum uint16
}

func NewClientMsg(commandType uint8, stationNum uint16) ClientMsg {
	CLIMsg := ClientMsg{
		MsgType:    commandType,
		StationNum: stationNum,
	}
	return CLIMsg
}

func CopyBytes(bytes []byte) []byte {
	n := len(bytes)
	newBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		newBytes[i] = bytes[i]
	}
	return newBytes
}
