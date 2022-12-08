package control

import (
	"context"
	"fmt"
	"log"
	"os"
	"snowcast/pkg/protocol"
	"snowcast/pkg/protocol/rpcMsg"
	"strconv"
)

func (cc *ClientControl) CallHelloRPC() {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	udpPort, err := strconv.Atoi(cc.UDPPort)
	if err != nil {
		log.Fatalln(err)
	}

	helloRequest := rpcMsg.RequestHello{
		MsgType: uint32(rpcMsg.RPCType_REQUEST_HELLO),
		UdpPort: uint32(udpPort),
		Address: ToIpColonPortNum("localhost", cc.UDPPort),
	}
	response, err := cc.ClientRPCClient.HandleHelloMsg(context.Background(), &helloRequest)
	if err != nil {
		log.Fatalf("Fail to send Hello RPC: %v\n", err)
	}
	fmt.Printf("there are %v songs in total\n", response.SongNum)
}

func (cc *ClientControl) HandleClientMsg() {
	for {
		clientCLI := <-cc.ClientMsgChan
		switch clientCLI.MsgType {
		case protocol.TypeSetStation:
			fmt.Println("hello, try to set station")
			// clientControl.SendSetStationMsg(clientCLI.StationNum)
		case protocol.TypeQuitClient:
			os.Exit(0)
		}
	}
}
