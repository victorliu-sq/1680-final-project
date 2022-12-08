package serverPkg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"snowcast/pkg/protocol/CLI"
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
	ServerCLIChan chan CLI.ServerCLIMsg
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
	server.ServerCLIChan = make(chan CLI.ServerCLIMsg, 1)
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

	// Open an listener
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
	}

	// broadcast udp data
	go server.InitialDaemonStations()

	// log.Println("Server is running successfully!")
	// Handle hello msg
	go server.RevClientConn()
	// Handle server CLI
	go server.ScanServerCLI()
	// go server.ScanServerCLI(server.ServerCLIChan)
	// Handle Server CLI Msg
}

// **********************************************************
// Initialization

func (server *Server) ToChunks(filename string) [][]byte {
	// fileName := "./mp3/Beethoven-SymphonyNo5.mp3"
	filePath := fmt.Sprintf("./mp3/%v", filename)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	// fio, err := os.Stat(filePath)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // fmt.Printf("file size is: %v\n", fio.Size())
	// chunksNum := (fio.Size() + maxPacketSize - 1) / maxPacketSize
	// fmt.Printf("chunks num is: %v\n", chunksNum)
	// read files into chunks
	chunks := [][]byte{}
	r := bufio.NewReader(f)
	for {
		chunk := make([]byte, server.MaxPacketSize)
		n, err := r.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		chunk = chunk[:n]
		chunks = append(chunks, chunk)
		// fmt.Println(chunk)
	}
	return chunks
}

// ************************************************************
// Port Conversion
func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}
