package server

import (
	"fmt"
	"net"
	"snowcast/pkg/protocol"
)

type Server struct {
	MaxPacketSize int
	// client
	// Listener *net.TCPListener
	// file
	Filenames []string
	// Filename2Chunks     map[string][][]byte
	StationIdx2Filename map[uint16]string
	// cli
	ServerMsgChan chan protocol.ServerMsg
	// control
	FirstHello      map[string]int
	FirstSetStation map[string]int
	// Controls            map[string]int
	Control2Conn        map[string]net.Conn
	Control2StationIdx  map[string]uint16
	Control2Listener    map[string]string
	StationIdx2Controls map[uint16]map[string]int

	// listener
	Listener2Conn map[string]*net.UDPConn
}

func (server *Server) Make(args []string) {
	server.MaxPacketSize = 4096
	server.Filenames = []string{}
	server.StationIdx2Filename = map[uint16]string{}
	server.StationIdx2Controls = map[uint16]map[string]int{}
	// server.Controls = map[string]int{}
	server.Control2StationIdx = map[string]uint16{}
	server.ServerMsgChan = make(chan protocol.ServerMsg, 1)
	server.Control2Listener = map[string]string{}
	server.Listener2Conn = map[string]*net.UDPConn{}
	// server.Filename2Chunks = map[string][][]byte{}
	server.Control2Conn = map[string]net.Conn{}
	server.FirstHello = map[string]int{}
	server.FirstSetStation = map[string]int{}

	for i := 2; i < len(args); i++ {
		idx := uint16(i - 2)
		filename := args[i]
		server.Filenames = append(server.Filenames, filename)
		server.StationIdx2Filename[idx] = args[i]
		server.StationIdx2Controls[idx] = map[string]int{}
		// file to chunks
		// We do not need to get chunks of files at first --> time consuming
		// chunks := server.ToChunks(filename)
		// server.Filename2Chunks[filename] = chunks
	}

	/* // Open an listener
	listenPort := args[1]
	// 1. specify a server port Number to get an TCP addr
	addr, err := net.ResolveTCPAddr("tcp4", ToColonPortNumber(listenPort))
	if err != nil {
		log.Fatalln(err)
	}

	// 2. create a listener listening on that TCP addr
	server.Listener, err = net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	} */
	// Handle server CLI Msg
	go server.ScanServerCLI()

	// broadcast udp data
	// go server.InitialDaemonStations()

	// log.Println("Server is running successfully!")
	// Handle hello msg
	// go server.RevClientConn()

	server.HandleServerMsg()
}

// ************************************************************
// Port Conversion
func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}
