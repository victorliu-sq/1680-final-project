package control

import (
	"context"
	"fmt"
	"log"
	"os"
	"snowcast/pkg/protocol"
	"snowcast/pkg/protocol/rpcMsg"
)

func (cc *ClientControl) HandleClientMsg() {
	for {
		clientCLI := <-cc.ClientMsgChan
		switch clientCLI.MsgType {
		case protocol.TypeSetStation:
			fmt.Println("hello, try to set station")
			cc.CallSetStationRPC(clientCLI.StationNum)
		case protocol.TypeQuitClient:
			os.Exit(0)
		}
	}
}

func (cc *ClientControl) CallHelloRPC() {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	request := rpcMsg.RequestHello{
		MsgType:     uint32(rpcMsg.RPCType_REQUEST_HELLO),
		UdpPort:     cc.UDPPort,
		ControlName: cc.ControlName,
	}
	response, err := cc.ClientRPCClient.HandleHelloMsg(context.Background(), &request)
	if err != nil {
		log.Fatalf("Fail to send Hello RPC: %v\n", err)
	}
	cc.IsWelcomeRev = true
	fmt.Printf("Welcome to Snowcast! The server has %v stations\n", response.SongNum)
	// fmt.Printf("there are %v songs in total\n", response.SongNum)
}

func (cc *ClientControl) CallSetStationRPC(stationNum uint16) {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	request := rpcMsg.RequestSetStation{
		MsgType:     uint32(rpcMsg.RPCType_REQUEST_SETSTATION),
		StationNum:  uint32(stationNum),
		ControlName: cc.ControlName,
	}
	response, err := cc.ClientRPCClient.HandleSetStationMsg(context.Background(), &request)
	if err != nil {
		log.Fatalf("Fail to send SetStation RPC: %v\n", err)
	}
	if response.MsgType == uint32(rpcMsg.RPCType_RESPOSNE_INVALID) {
		fmt.Println("invalid setStation Msg")
		return
	}
	fmt.Printf("New song announced: %v\n", response.SongName)
}
