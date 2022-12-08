package main

import (
	"log"
	"os"
	"snowcast/pkg/protocol/CLI"
	"snowcast/pkg/serverPkg"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Usage %v <server listen port> + <filename> ...", os.Args[0])
	}
	server := serverPkg.Server{}
	server.Make(os.Args)

	// Rev and Handle serverCLI
	for {
		serverCLIMsg := <-server.ServerCLIChan
		switch serverCLIMsg.CLIType {
		case CLI.TypePrintClients:
			// fmt.Println(server.StationIdx2Controls)
			server.HandleCLIPrintClients()
		case CLI.TypeQuitServer:
			server.HandleCLIQuitServer()
			// os.Exit(0) still works in goroutine
		case CLI.TypeRemoveStationIdx:
			server.HandleCLIRemoveStationIdx(serverCLIMsg)
		case CLI.TypeAddFile:
			server.HandleCLIAddFile(serverCLIMsg)
		case CLI.TypePrintToFile:
			server.HandleCLIPrintToFile(serverCLIMsg)
		}
	}
}
