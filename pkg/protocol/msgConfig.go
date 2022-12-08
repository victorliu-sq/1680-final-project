package protocol

const (
	// server cli
	TypePrintClients     = uint8(0)
	TypeQuitServer       = uint8(1)
	TypeRemoveStationIdx = uint8(2)
	TypePrintToFile      = uint8(4)
	// client cli
	TypeQuitClient = uint8(5)
	// client command
	TypeHello      = uint8(6)
	TypeSetStation = uint8(7)
	TypeUnknown    = uint8(100)
)
