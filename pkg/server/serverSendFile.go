package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func (server *Server) InitialDaemonStations() {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	for idx, filename := range server.Filenames {
		go server.DaemonStation(idx, filename)
	}
}

func (server *Server) DaemonStation(stationIdx int, filename string) {
	// fmt.Printf("Start a Daemon for station %v to send chunks to its listeners\n", stationIdx)
	// Open a file
	// fileName := "./mp3/Beethoven-SymphonyNo5.mp3"
	// filePath := fmt.Sprintf("./mp3/%v", filename)
	// f, err := os.Open(filePath)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	// Start to send Announce Msg to controller
	// server.BroadcastToControllers(stationIdx)
	// Get a reader
	r := bufio.NewReader(f)
	// while this stationIdx exists in server, the daemon should broadcast to listeners
	for server.CheckStationIdxInServerX(stationIdx) {
		// Read one chunk from the file
		chunk := make([]byte, server.MaxPacketSize)
		n, err := r.Read(chunk)
		if err != nil {
			if err == io.EOF {
				// reset the reader
				_, err := f.Seek(0, io.SeekStart)
				if err != nil {
					log.Fatalln(err)
				}
				// Start again from file start, we need to send announceMsg to all clients
				// fmt.Printf("Broadcast stationIdx %v Announcement to controllers\n", stationIdx)
				// server.BroadcastToControllers(stationIdx)
				continue
			}
			log.Fatalln(err)
		}
		chunk = chunk[:n]
		go server.BroadcastToListeners(stationIdx, chunk)
		// 16KiB/s
		// = 2 ^ 4 * 2 ^ 10 B /s
		// = 2 ^ 14 B/s
		// if one chunk has 4096 (2 ^ 12) bytes
		// = 2 ^ 2 chunk/s

		// time.Sleep(1 * time.Second)
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Printf("Close the a Daemon for station %v to send chunks\n", stationIdx)
}

func (server *Server) BroadcastToListeners(stationIdx int, chunk []byte) {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	controlAddrs := server.StationIdx2Controls[uint16(stationIdx)]
	for controlAddr, _ := range controlAddrs {
		listenerAddr := server.Control2Listener[controlAddr]
		conn := server.Listener2Conn[listenerAddr]
		go server.SendToListner(conn, chunk)
	}
}

func (server *Server) SendToListner(conn *net.UDPConn, chunk []byte) {
	conn.Write(chunk)
	// _, err := conn.Write(chunk)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Printf("%v:%v\n", controlAddr, listenerAddr)
}

// Broadcast to controllers
// func (server *Server) BroadcastToControllers(stationIdx int) {
// 	server.Mu.Lock()
// 	defer server.Mu.Unlock()
// 	filename := server.Filenames[stationIdx]
// 	controlAddrs := server.StationIdx2Controls[uint16(stationIdx)]
// 	for controlAddr, _ := range controlAddrs {
// 		// fmt.Println(controlAddr)
// 		go server.SendAnnounceMsgX(controlAddr, filename)
// 		// listenerAddr := server.Control2Listener[controlAddr]
// 		// conn := server.Listener2Conn[listenerAddr]
// 		// go server.SendToListner(conn, chunk)
// 	}
// }

func (server *Server) CheckStationIdxInServerX(stationIdx int) bool {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	if _, ok := server.StationIdx2Filename[uint16(stationIdx)]; ok {
		return true
	}
	return false
}
