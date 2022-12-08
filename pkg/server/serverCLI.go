package server

import (
	"bufio"
	"os"
	"snowcast/pkg/protocol"
)

func (server *Server) ScanServerCLI() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 1 && line[0] == 'p' {
				serverCLIMsg := protocol.NewServerCLI(protocol.TypePrintClients, 0, "")
				server.ServerMsgChan <- serverCLIMsg

			} else if len(line) == 1 && line[0] == 'q' {
				serverCLIMsg := protocol.NewServerCLI(protocol.TypeQuitServer, 0, "")
				server.ServerMsgChan <- serverCLIMsg
			} else if line[0] == 'p' {
				if len(line) < 3 {
					continue
				}
				filename := line[2:]
				// fmt.Printf("Try to add one station for file %v\n", filename)
				serverCLIMsg := protocol.NewServerCLI(protocol.TypePrintToFile, 0, filename)
				server.ServerMsgChan <- serverCLIMsg
			}
		}
	}
}
