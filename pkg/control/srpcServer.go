package control

import (
	"context"
	"fmt"
	"log"
	"net"
	"snowcast/pkg/protocol/rpcMsg"

	"google.golang.org/grpc"
)

type SRPCServer struct {
	rpcMsg.ServerMsgServiceServer
	control *ClientControl // snowcast server
}

func (srpcServer *SRPCServer) Make(cc *ClientControl) {
	srpcServer.control = cc
}

func (cc *ClientControl) OpenSRPCServer(listenPort string) {
	// 1. specify a server port Number to get a listener
	listener, err := net.Listen("tcp", GetListenerAddress(listenPort))
	if err != nil {
		log.Fatalln(err)
	}

	srpcServer := SRPCServer{}
	srpcServer.Make(cc)

	// 2. open a gRPC server
	grpcServer := grpc.NewServer()

	rpcMsg.RegisterServerMsgServiceServer(grpcServer, &srpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(srpcServer)
}

// func (srpcServer *SRPCServer) HandleSendFileMsg(ctx context.Context, request *rpcMsg.ResponseAnnounce) (*rpcMsg.ResponseAnnounce, error) {
// 	fmt.Printf("Announce songname :%v", request.SongName)
// 	return &rpcMsg.ResponseAnnounce{}, nil
// }

func (srpcServer *SRPCServer) HandleSendFileMsg(ctx context.Context, request *rpcMsg.RequestSendFile) (*rpcMsg.ResponseSendFile, error) {
	fmt.Printf("Get Announce(Sendfile) %v \n", request.SongName)
	return &rpcMsg.ResponseSendFile{}, nil
}

func (srpcServer *SRPCServer) HandleShutdownMsg(ctx context.Context, request *rpcMsg.RequestShutdown) (*rpcMsg.ResponseShutdown, error) {
	fmt.Printf("station %v has been shutdown, please choose another one\n", request.StationNum)
	return &rpcMsg.ResponseShutdown{}, nil
}
