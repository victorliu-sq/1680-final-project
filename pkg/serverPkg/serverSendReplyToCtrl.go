package serverPkg

import (
	"fmt"
	"log"
	"net"
	"snowcast/pkg/protocol/reply"
)

func (server *Server) SendWelcomeMsg(conn net.Conn, wMsg *reply.WelcomeMsg) {
	bytes := wMsg.Marsharl()
	// conn.Write(bytes)
	cur, total := 0, len(bytes)
	for cur < total {
		n, err := conn.Write(bytes[cur:])
		if err != nil {
			log.Fatalln(err)
		}
		cur += n
	}
}

func (server *Server) SendAnnounceMsg(conn net.Conn, songnameS string) {
	// log.Println("Try to send an announce Msg successfully")
	announceMsg := reply.NewAnnounceMsg(songnameS)
	bytes := announceMsg.Marsharl()
	// conn.Write(bytes)
	cur, total := 0, len(bytes)
	for cur < total {
		n, err := conn.Write(bytes[cur:])
		if err != nil {
			log.Fatalln(err)
		}
		cur += n
	}
	// conn.Write(bytes[:1])
	// time.Sleep(10 * time.Millisecond)
	// conn.Write(bytes[1:])
	log.Println("Send an announce Msg successfully")
}

func (server *Server) SendInvalidMsg(conn net.Conn, replyString string) {
	// log.Println("Try to send an Invalid reply msg for hello command")
	invalidMsg := reply.NewInvalidCommandMsg(replyString)
	bytes := invalidMsg.Marsharl()
	// conn.Write(bytes)
	cur, total := 0, len(bytes)
	for cur < total {
		n, err := conn.Write(bytes[cur:])
		if err != nil {
			log.Fatalln(err)
		}
		cur += n
	}
	log.Println("Send an Invalid Reply Msg successfully")
}

// Broadcast NewStation Msg
func (server *Server) BroadcastNewStationMsg(stationNum uint16) {
	nMsg := reply.NewNewStationMsg(stationNum)
	bytes := nMsg.Marsharl()
	for _, conn := range server.Control2Conn {
		conn.Write(bytes)
		fmt.Println("Send NewStationMsg back")
	}
}

// Broadcast StationShutDown Msg
func (server *Server) BroadcasStationShutDownMsg(stationIdx uint16) {
	nMsg := reply.NewStationShutDown(stationIdx)
	bytes := nMsg.Marsharl()
	for controlAddr, _ := range server.StationIdx2Controls[stationIdx] {
		conn := server.Control2Conn[controlAddr]
		conn.Write(bytes)
		fmt.Println("Send StationShutDownMsg back")
	}
}

func (server *Server) SendAnnounceMsgX(controlAddr, songnameS string) {
	// log.Println("Try to send an announce Msg successfully")
	server.Mu.Lock()
	defer server.Mu.Unlock()
	conn, ok := server.Control2Conn[controlAddr]
	if !ok {
		return
	}
	announceMsg := reply.NewAnnounceMsg(songnameS)
	bytes := announceMsg.Marsharl()
	// conn.Write(bytes)
	cur, total := 0, len(bytes)
	for cur < total {
		n, err := conn.Write(bytes[cur:])
		if err != nil {
			log.Fatalln(err)
		}
		cur += n
	}
	// log.Println("Send an announce Msg successfully")
}
