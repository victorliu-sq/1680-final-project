package server

import (
	"fmt"
	"net"
	"snowcast/pkg/protocol"
	"snowcast/pkg/protocol/rpcMsg"
	"sync"
)

type Server struct {
	MaxPacketSize int
	// Sync
	Mu sync.Mutex
	// client
	Listener *net.TCPListener
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
	Control2SRPCClient  map[string]rpcMsg.ServerMsgServiceClient
	Control2Conn        map[string]net.Conn
	Control2StationIdx  map[string]uint16
	Control2Listener    map[string]string
	StationIdx2Controls map[uint16]map[string]int

	// listener
	Listener2Conn map[string]*net.UDPConn
}

func (server *Server) Make(args []string) {
	server.Mu = sync.Mutex{}
	server.MaxPacketSize = 4096
	server.Filenames = []string{}
	server.StationIdx2Filename = map[uint16]string{}
	server.StationIdx2Controls = map[uint16]map[string]int{}
	// server.Controls = map[string]int{}
	server.Control2StationIdx = map[string]uint16{}
	server.ServerMsgChan = make(chan protocol.ServerMsg, 1)
	server.Control2Listener = map[string]string{}
	server.Listener2Conn = map[string]*net.UDPConn{}
	server.Control2SRPCClient = map[string]rpcMsg.ServerMsgServiceClient{}
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
	}

	// Handle server CLI Msg
	go server.ScanServerCLI()

	// broadcast udp data
	go server.InitialDaemonStations()

	// log.Println("Server is running successfully!")
	go server.HandleServerCLI()

	// Open an listener for RPC Msg
	listenPort := args[1]
	server.OpenCRPCServer(listenPort)
}

// ************************************************************
// Port Conversion
func GetListenerAddress(portNum string) string {
	return fmt.Sprintf("localhost:%v", portNum)
}
