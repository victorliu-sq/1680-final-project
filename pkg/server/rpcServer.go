package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"snowcast/pkg/protocol/rpcMsg"
	"strconv"
	"strings"
)

type RPCServer struct {
	rpcMsg.ControlMsgServiceServer
	snowcastServer *Server // snowcast server
}

func (rpcServer *RPCServer) Make(snowcastServer *Server) {
	rpcServer.snowcastServer = snowcastServer
}

// Server RPC Handlers
// **************************************************************************
func (rpcServer *RPCServer) HandleHelloMsg(ctx context.Context, request *rpcMsg.RequestHello) (*rpcMsg.ResponseWelcome, error) {
	rpcServer.snowcastServer.Mu.Lock()
	defer rpcServer.snowcastServer.Mu.Unlock()
	server := rpcServer.snowcastServer
	fmt.Printf("receive one Hello Msg with port %v\n", request.UdpPort)

	// Dial the udp port and start a goroutine to send mp3 data
	listenerAddr := ToIpColonPortNum("localhost", request.UdpPort)
	// listenerAddr := GetUdpAddr(controlAddr, helloMsg.UdpPort)
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

	// update UDP metadata
	server.Control2Listener[request.ControlName] = listenerAddr

	// return back responses
	return &rpcMsg.ResponseWelcome{
		MsgType: uint32(rpcMsg.RPCType_RESPONSE_WELCOME),
		SongNum: uint32(len(rpcServer.snowcastServer.Filenames)),
	}, nil
}

// helper functions
// **************************************************************************
func GetUdpAddr(controlAddr string, udpPort uint16) string {
	ipAddr := strings.Split(controlAddr, ":")[0]
	return fmt.Sprintf("%v:%v", ipAddr, strconv.Itoa(int(udpPort)))
}

func ToIpColonPortNum(ipAddr, portNum string) string {
	return fmt.Sprintf("%s:%s", ipAddr, portNum)
}

func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}
