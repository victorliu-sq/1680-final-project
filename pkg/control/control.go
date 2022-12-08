package control

import (
	"fmt"
	"log"
	"snowcast/pkg/protocol"
	"snowcast/pkg/protocol/rpcMsg"
	"sync"

	"google.golang.org/grpc"
)

type ClientControl struct {
	Mu            sync.Mutex
	UDPPort       string
	ServerIP      string
	ServerPort    string
	ControlName   string
	ClientMsgChan chan protocol.ClientMsg
	// ServerConn    *net.TCPConn
	ClientRPCClient rpcMsg.ControlMsgServiceClient
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
	cc.ClientMsgChan = make(chan protocol.ClientMsg, 1)
	// Dial to server
	// addr, err := net.ResolveTCPAddr("tcp4", ToIpColonPortNum(cc.ServerIP, cc.ServerPort))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // 2. Connect to server and get socket
	// conn, err = net.DialTCP("tcp4", nil, addr)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	conn, err := grpc.Dial(ToColonPortNumber(cc.ServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect: %v \n", err)
	}
	defer conn.Close()

	// client RPC client
	cc.ClientRPCClient = rpcMsg.NewControlMsgServiceClient(conn)

	// test for multiple hello msgs for an invalid reply
	go cc.CallHelloRPC()

	// test for sending an msg of unknown type
	// SendUnknownMsg(conn, portNum)
	go cc.ScanClientCLI()
	cc.HandleClientMsg()
}

// helper function
func ToIpColonPortNum(ipAddr, portNum string) string {
	return fmt.Sprintf("%s:%s", ipAddr, portNum)
}

func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}
