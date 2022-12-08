package serverPkg

import (
	"fmt"
	"log"
	"net"
	"snowcast/pkg/protocol/CLI"
	"snowcast/pkg/protocol/command"
	"time"
)

func (server *Server) RevClientConn() {
	for {
		// 3. Periodically receive connection from client and get socket
		// conn, err := listener.Accept()
		conn, err := server.Listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		addr := conn.RemoteAddr().String()
		server.Mu.Lock()
		server.Control2Conn[addr] = conn
		server.Mu.Unlock()
		// log.Printf("Server gets one connection from %v!\n", addr)
		fmt.Printf("Server gets one connection from Control %v!\n", addr)
		// start a goroutine to handle all commands from 1 client
		go server.RevCommandMsg(conn, addr)
	}
}

func (server *Server) RevCommandMsg(conn net.Conn, addr string) {
	defer server.HandleDisconnection(addr)
	defer conn.Close()
	waitForHello := true
	for {
		if waitForHello {
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			waitForHello = false
		} else {
			conn.SetReadDeadline(time.Now().Add(3000 * time.Second))
		}
		// 4. Read bytes from socket
		bytes := make([]byte, 3)
		bnum, err := conn.Read(bytes)
		bytes = bytes[:bnum]
		// fmt.Println(bnum)
		// If the client disconnects, the bnum will be 0
		if bnum < 1 || err != nil {
			// fmt.Printf("Client %v lost bits", conn.RemoteAddr().String())
			log.Printf("Client %v disconnects\n", addr)
			// server.HandleDisconnection(addr)
			return
		}
		// Not enough bytes are read
		cur, total := bnum, 3
		for cur < total {
			conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			temp := make([]byte, 3)
			bnum, err = conn.Read(temp)
			if err != nil {
				return
			}
			temp = temp[:bnum]
			bytes = append(bytes, temp...)
			cur += bnum
		}
		commandType := uint8(bytes[0])
		// fmt.Printf("command type is %v\n", commandType)
		switch commandType {
		case command.TypeHello:
			if !server.HandleHelloMsg(conn, addr, CLI.CopyBytes(bytes)) {
				return
			}
		case command.TypeSetStation:
			if !server.HandleSetStationMsg(conn, CLI.CopyBytes(bytes)) {
				return
			}
		default:
			server.SendInvalidMsg(conn, "Invalid command type")
			return
		}
	}
}

func (server *Server) HandleDisconnection(addr string) {
	// When a client disconnects from a server
	// remove it from server.ControlAddr2StationIdx
	server.Mu.Lock()
	defer server.Mu.Unlock()
	stationIdx := server.Control2StationIdx[addr]
	delete(server.Control2Conn, addr)
	delete(server.Control2StationIdx, addr)
	// remove it from server.StationIdx2Controls
	// fmt.Println(server.StationIdx2Controls)
	delete(server.StationIdx2Controls[stationIdx], addr)
	// fmt.Println(server.StationIdx2Controls)
	// remove it from first Hello and first SetStation
	delete(server.FirstHello, addr)
	delete(server.FirstSetStation, addr)
	// fmt.Printf("Handle disconnection %v\n", addr)
}
