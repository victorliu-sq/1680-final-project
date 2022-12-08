package serverPkg

import (
	"fmt"
	"log"
	"net"
	"snowcast/pkg/protocol/command"
	"snowcast/pkg/protocol/reply"
	"strconv"
	"strings"
)

func (server *Server) HandleHelloMsg(conn net.Conn, controlAddr string, bytes []byte) bool {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	// Check if this is 1st hello msg of this client
	if _, ok := server.FirstHello[controlAddr]; ok {
		server.SendInvalidMsg(conn, "The Client sends more than one Hello Msg\n")
		return false
	}
	// Dial the udp port and start a goroutine to send mp3 data
	helloMsg := command.UnmarsharlHelloMsg(bytes)
	listenerAddr := GetUdpAddr(controlAddr, helloMsg.UdpPort)
	// fmt.Println(listenerAddr)

	udpAddr, err := net.ResolveUDPAddr("udp", listenerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	server.Listener2Conn[listenerAddr] = udpConn
	// update numStations & clientSetHello
	server.FirstHello[controlAddr] = 1
	// update UDP metadata
	server.Control2Listener[controlAddr] = listenerAddr

	// fmt.Println("control2listener:", control2listener)
	// fmt.Println("listener2conn:", listener2conn)
	// send welcome msg back to server
	wMsg := reply.NewWelcomeMsg(uint16(len(server.Filenames)))
	// fmt.Printf("Receive one Hello Msg from %v\n", controlAddr)
	server.SendWelcomeMsg(conn, wMsg)
	fmt.Printf("Send one Welcome Msg to control %v successfully\n", controlAddr)
	return true
}

func (server *Server) HandleSetStationMsg(conn net.Conn, setStationBytes []byte) bool {
	// fmt.Println("Receive one SetStation Msg")
	server.Mu.Lock()
	defer server.Mu.Unlock()
	// Dial the Udp port
	// fmt.Println(server.StationIdx2Controls)
	setStationMsg := command.UnmarshalSetstationMsg(setStationBytes)
	stationIdx := setStationMsg.StationNumber
	controlAddr := conn.RemoteAddr().String()
	// Check if stationIdx is out of range
	total := uint16(len(server.StationIdx2Filename))
	if stationIdx >= total {
		server.SendInvalidMsg(conn, "The Client's stationIdx is out of range\n")
		return false
	}
	// Check if setStation Msg is sent before Hello Msg
	if _, ok := server.FirstHello[controlAddr]; !ok {
		server.SendInvalidMsg(conn, "Send SetStation Msg before Hello Msg\n")
		return false
	}
	// Update metadata
	if _, ok := server.FirstSetStation[controlAddr]; !ok {
		// if this is the first setStationMsg
		server.FirstSetStation[controlAddr] = 1
		server.StationIdx2Controls[stationIdx][controlAddr] = 1
		server.Control2StationIdx[controlAddr] = stationIdx
	} else {
		// if this is not the first setStationMsg
		prevStationIdx := server.Control2StationIdx[controlAddr]
		// update client's stationIdx
		if prevStationIdx != stationIdx {
			server.Control2StationIdx[controlAddr] = stationIdx
		}
		// remove client from prevStation's client set
		delete(server.StationIdx2Controls[prevStationIdx], controlAddr)
		// add client to newStation's client set
		server.StationIdx2Controls[stationIdx][controlAddr] = 1
	}
	// Send back AnnounceMsg
	songnameS := server.StationIdx2Filename[stationIdx]
	server.SendAnnounceMsg(conn, songnameS)
	return true
}

func GetUdpAddr(controlAddr string, udpPort uint16) string {
	ipAddr := strings.Split(controlAddr, ":")[0]
	return fmt.Sprintf("%v:%v", ipAddr, strconv.Itoa(int(udpPort)))
}
