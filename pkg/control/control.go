package control

import (
	"fmt"
	"snowcast/pkg/protocol"
	"snowcast/pkg/protocol/rpcMsg"
	"sync"
)

type ClientControl struct {
	Mu             sync.Mutex
	UDPPort        string
	ServerIP       string
	ServerPort     string
	ControlName    string
	SRPCServerPort string
	ClientMsgChan  chan protocol.ClientMsg

	StationNum uint16
	// ServerConn    *net.TCPConn
	CRPCClient rpcMsg.ControlMsgServiceClient
	// Check whether the server sends a Welcome before client sends a Hello
	IsHelloSent bool
	// Check whether server has received one Welcome Msg
	IsWelcomeRev bool
	//	Check whether the server sends an Announce before the client has sent a SetStation
	IsSetStationSent bool
}

func (cc *ClientControl) Make(args []string) {
	cc.Mu = sync.Mutex{}
	cc.ServerIP, cc.ServerPort, cc.UDPPort = args[1], args[2], args[3]
	cc.ControlName = args[4]
	cc.SRPCServerPort = args[5]
	cc.ClientMsgChan = make(chan protocol.ClientMsg, 1)
	// Dial to server
	cc.CreateCRPCClient()

	// test for multiple hello msgs for an invalid reply
	go cc.CallHelloRPC()

	// test for sending an msg of unknown type
	// SendUnknownMsg(conn, portNum)
	go cc.ScanClientCLI()
	go cc.HandleClientMsg()
	// receive sendFileRequests and shutDownRequests
	cc.OpenSRPCServer(cc.SRPCServerPort)

}

// helper function
// ***************************************************************************
func ToIpColonPortNum(ipAddr, portNum string) string {
	return fmt.Sprintf("%s:%s", ipAddr, portNum)
}

func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}

func GetListenerAddress(portNum string) string {
	return fmt.Sprintf("localhost:%v", portNum)
}
