package server

import (
	"bufio"
	"fmt"
	"os"
	"snowcast/pkg/protocol"
	"strconv"
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
			} else if line[0] == 'r' {
				// check length of "r + stationIdx"
				if len(line) < 3 {
					continue
				}
				// check whether or not the stationIdx exists in the server
				if _, err := strconv.Atoi(line[2:]); err != nil {
					continue
				}
				idx, _ := strconv.Atoi(line[2:])
				// check whether stationIdx exists
				if !server.CheckStationIdxInServerX(idx) {
					fmt.Printf("The stationIdx %v to remove does not exists in Server\n", idx)
					continue
				}
				// fmt.Printf("Try to remove one station %v\n", idx)
				serverCLIMsg := protocol.NewServerCLI(protocol.TypeRemoveStationIdx, uint16(idx), "")
				server.ServerMsgChan <- serverCLIMsg
			}
		}
	}
}
