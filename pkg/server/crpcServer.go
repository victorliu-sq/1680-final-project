package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"snowcast/pkg/protocol/rpcMsg"

	"google.golang.org/grpc"
)

type CRPCServer struct {
	rpcMsg.ControlMsgServiceServer
	snowcastServer *Server // snowcast server
}

func (crpcServer *CRPCServer) Make(snowcastServer *Server) {
	crpcServer.snowcastServer = snowcastServer
}

func (server *Server) OpenCRPCServer(listenPort string) {
	// 1. specify a server port Number to get a listener
	listener, err := net.Listen("tcp", GetListenerAddress(listenPort))
	if err != nil {
		log.Fatalln(err)
	}

	crpcServer := CRPCServer{}
	crpcServer.Make(server)

	// 2. open a gRPC server
	grpcServer := grpc.NewServer()

	rpcMsg.RegisterControlMsgServiceServer(grpcServer, &crpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}

// Server RPC Handlers
// **************************************************************************
func (crpcServer *CRPCServer) HandleHelloMsg(ctx context.Context, request *rpcMsg.RequestHello) (*rpcMsg.ResponseWelcome, error) {
	crpcServer.snowcastServer.Mu.Lock()
	defer crpcServer.snowcastServer.Mu.Unlock()
	server := crpcServer.snowcastServer
	controlAddr := request.ControlName
	srpcServerPort := request.SrpcServerPort
	fmt.Printf("receive one Hello Msg with port %v\n", request.UdpPort)
	// Check if this is 1st hello msg of this client
	if _, ok := server.FirstHello[controlAddr]; ok {
		return &rpcMsg.ResponseWelcome{
			MsgType: uint32(rpcMsg.RPCType_RESPOSNE_INVALID),
		}, nil
	}
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
	server.FirstHello[controlAddr] = 1
	server.Control2Listener[controlAddr] = listenerAddr

	// create corresponding Server RPC Service Client
	server.CreateSRPCClient(controlAddr, srpcServerPort)
	// fmt.Println(server.Control2SRPCClient)
	// return back responses
	return &rpcMsg.ResponseWelcome{
		MsgType: uint32(rpcMsg.RPCType_RESPONSE_WELCOME),
		SongNum: uint32(len(crpcServer.snowcastServer.Filenames)),
	}, nil
}

func (crpcServer *CRPCServer) HandleSetStationMsg(ctx context.Context, request *rpcMsg.RequestSetStation) (*rpcMsg.ResponseAnnounce, error) {
	server := crpcServer.snowcastServer
	stationIdx := uint16(request.StationNum)
	controlAddr := request.ControlName

	// Check if stationIdx is out of range
	total := uint16(len(server.StationIdx2Filename))
	if stationIdx >= total {
		return &rpcMsg.ResponseAnnounce{
			MsgType: uint32(rpcMsg.RPCType_RESPOSNE_INVALID),
		}, nil
	}
	// Check if setStation Msg is sent before Hello Msg
	if _, ok := server.FirstHello[controlAddr]; !ok {
		fmt.Println(ok)
		return &rpcMsg.ResponseAnnounce{
			MsgType: uint32(rpcMsg.RPCType_RESPOSNE_INVALID),
		}, nil
	}
	// Update metadata
	if _, ok := server.FirstSetStation[controlAddr]; !ok {
		// if this is the first setStationMsg
		server.FirstSetStation[controlAddr] = 1
		server.StationIdx2Controls[stationIdx][controlAddr] = 1
		server.Control2StationIdx[controlAddr] = stationIdx
	} else {
		// if this is not the first setStationMsg
		prevStationIdx := server.Control2StationIdx[controlAddr]
		// update client's stationIdx
		if prevStationIdx != stationIdx {
			server.Control2StationIdx[controlAddr] = stationIdx
		}
		// remove client from prevStation's client set
		delete(server.StationIdx2Controls[prevStationIdx], controlAddr)
		// add client to newStation's client set
		server.StationIdx2Controls[stationIdx][controlAddr] = 1
	}

	// return response
	songName := server.StationIdx2Filename[uint16(stationIdx)]
	return &rpcMsg.ResponseAnnounce{
		MsgType:  uint32(rpcMsg.RPCType_RESPONSE_ANNOUNCE),
		SongName: songName,
	}, nil
}

// helper functions
// *************************************************************************

func ToIpColonPortNum(ipAddr, portNum string) string {
	return fmt.Sprintf("%s:%s", ipAddr, portNum)
}

func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}
