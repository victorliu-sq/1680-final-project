package protocol

const (
	// server cli
	TypePrintClients     = uint8(0)
	TypeQuitServer       = uint8(1)
	TypeRemoveStationIdx = uint8(2)
	TypePrintToFile      = uint8(4)
	// server reply
	TypeWelcome         = uint8(5)
	TypeAnnounce        = uint8(6)
	TypeInvalid         = uint8(7)
	TypeNewStation      = uint8(8)
	TypeStationShutDown = uint8(9)
	// client cli
	TypeQuitClient = uint8(10)
	// client command
	TypeHello      = uint8(11)
	TypeSetStation = uint8(12)
)
