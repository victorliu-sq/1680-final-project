package protocol

type ServerMsg struct {
	CLIType    uint8
	StationNum uint16
	Filename   string
}

func NewServerCLI(commandType uint8, stationNum uint16, filename string) ServerMsg {
	CLIMsg := ServerMsg{
		CLIType:    commandType,
		StationNum: stationNum,
		Filename:   filename,
	}
	return CLIMsg
}
