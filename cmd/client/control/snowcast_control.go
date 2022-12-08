package main

import (
	"log"
	"os"
	clientPkg "snowcast/pkg/controlPkg"
	"snowcast/pkg/protocol/CLI"
	"snowcast/pkg/protocol/command"
)

func main() {
	// check os.Args #
	if len(os.Args) < 4 {
		log.Fatalf("Usage:  %s <server IP> <server port> <udp port>", os.Args[0])
	}
	clientControl := clientPkg.ClientControl{}
	clientControl.Make(os.Args)
	defer clientControl.ServerConn.Close()

	// Receive Welcome reply + Announce reply + Invalid reply msg
	go clientControl.RevReplyMsg()
	// Scan Client Command from
	go clientControl.ScanClientCLI(clientControl.ClientCLIChan)
	// 3. Send Hello msg to snowcast
	go clientControl.SendHelloMsg(clientControl.UDPPort)

	// test for multiple hello msgs for an invalid reply
	// SendHelloMsg(conn, portNum)

	// test for sending an msg of unknown type
	// SendUnknownMsg(conn, portNum)
	// Deal with ClientCLI
	for {
		clientCLI := <-clientControl.ClientCLIChan
		switch clientCLI.CLIType {
		case command.TypeSetStation:
			clientControl.SendSetStationMsg(clientCLI.StationNum)
		case CLI.TypeQuitClient:
			os.Exit(0)
		}
	}
}
