package server

import (
	"log"
	"snowcast/pkg/protocol/rpcMsg"

	"google.golang.org/grpc"
)

func (server *Server) CreateSRPCClient(controlName, listenPort string) {
	conn, err := grpc.Dial(ToColonPortNumber(listenPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect: %v \n", err)
	}
	// client RPC client
	srpcClient := rpcMsg.NewServerMsgServiceClient(conn)
	server.Control2SRPCClient[controlName] = srpcClient
}
