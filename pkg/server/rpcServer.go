package server

import (
	"context"
	"fmt"
	"snowcast/pkg/protocol/rpcMsg"
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
	// server.Mu.Lock()
	// defer server.Mu.Unlock()
	fmt.Printf("receive one Hello Msg with port %v\n", request.UdpPort)
	return &rpcMsg.ResponseWelcome{
		MsgType: uint32(rpcMsg.RPCType_RESPONSE_WELCOME),
		SongNum: uint32(len(rpcServer.snowcastServer.Filenames)),
	}, nil
}
