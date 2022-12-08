package control

import (
	"context"
	"fmt"
	"log"
	"snowcast/pkg/protocol/rpcMsg"

	"google.golang.org/grpc"
)

func (cc *ClientControl) CreateCRPCClient() {
	conn, err := grpc.Dial(ToColonPortNumber(cc.ServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect: %v \n", err)
	}
	// defer conn.Close()
	// client RPC client
	cc.CRPCClient = rpcMsg.NewControlMsgServiceClient(conn)
}

func (cc *ClientControl) CallHelloRPC() {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	request := rpcMsg.RequestHello{
		MsgType:        uint32(rpcMsg.RPCType_REQUEST_HELLO),
		UdpPort:        cc.UDPPort,
		ControlName:    cc.ControlName,
		SrpcServerPort: cc.SRPCServerPort,
	}
	response, err := cc.CRPCClient.HandleHelloMsg(context.Background(), &request)
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
	if stationNum == cc.StationNum {
		fmt.Println("Same StationNum, please input a different one")
		return
	}
	request := rpcMsg.RequestSetStation{
		MsgType:     uint32(rpcMsg.RPCType_REQUEST_SETSTATION),
		StationNum:  uint32(stationNum),
		ControlName: cc.ControlName,
	}
	response, err := cc.CRPCClient.HandleSetStationMsg(context.Background(), &request)
	if err != nil {
		log.Fatalf("Fail to send SetStation RPC: %v\n", err)
	}
	if response.MsgType == uint32(rpcMsg.RPCType_RESPOSNE_INVALID) {
		fmt.Println("invalid setStation Msg")
		return
	}
	cc.StationNum = stationNum
	fmt.Printf("New song announced: %v\n", response.SongName)
}
