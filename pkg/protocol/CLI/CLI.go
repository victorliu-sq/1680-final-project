package CLI

const (
	// server cli
	TypePrintClients     = uint8(0)
	TypeQuitServer       = uint8(1)
	TypeRemoveStationIdx = uint8(2)
	TypeAddFile          = uint8(3)
	TypePrintToFile      = uint8(4)
	// client cli
	TypeQuitClient = uint8(4)
)

type ServerCLIMsg struct {
	CLIType    uint8
	StationNum uint16
	Filename   string
}

func NewServerCLI(commandType uint8, stationNum uint16, filename string) *ServerCLIMsg {
	CLIMsg := &ServerCLIMsg{
		CLIType:    commandType,
		StationNum: stationNum,
		Filename:   filename,
	}
	return CLIMsg
}

type ClientCLI struct {
	CLIType    uint8
	StationNum uint16
}

func NewClientCLI(commandType uint8, stationNum uint16) *ClientCLI {
	CLIMsg := &ClientCLI{
		CLIType:    commandType,
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
