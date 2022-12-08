package serverPkg

import (
	"bufio"
	"fmt"
	"os"
	"snowcast/pkg/protocol/CLI"
	"strconv"
)

func (server *Server) ScanServerCLI() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 1 && line[0] == 'p' {
				serverCLIMsg := CLI.NewServerCLI(CLI.TypePrintClients, 0, "")
				server.ServerCLIChan <- *serverCLIMsg

			} else if len(line) == 1 && line[0] == 'q' {
				serverCLIMsg := CLI.NewServerCLI(CLI.TypeQuitServer, 0, "")
				server.ServerCLIChan <- *serverCLIMsg
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
				serverCLIMsg := CLI.NewServerCLI(CLI.TypeRemoveStationIdx, uint16(idx), "")
				server.ServerCLIChan <- *serverCLIMsg
			} else if line[0] == 'a' {
				// if length
				if len(line) < 3 {
					continue
				}
				filename := line[2:]
				// check whether the file exists in map3 file folder
				if !server.CheckFileInFolder(filename) {
					fmt.Println("The file to add does not exist!")
					continue
				}
				// check whether the filename exists in server
				if server.CheckFileInServerX(filename) {
					fmt.Println("The file to add exists in server!")
					continue
				}
				// fmt.Printf("Try to add one station for file %v\n", filename)
				serverCLIMsg := CLI.NewServerCLI(CLI.TypeAddFile, 0, filename)
				server.ServerCLIChan <- *serverCLIMsg
			} else if line[0] == 'p' {
				if len(line) < 3 {
					continue
				}
				filename := line[2:]
				// fmt.Printf("Try to add one station for file %v\n", filename)
				serverCLIMsg := CLI.NewServerCLI(CLI.TypePrintToFile, 0, filename)
				server.ServerCLIChan <- *serverCLIMsg
			}
		}
	}
}

// ***************************************************************************
// Control helper function

func (server *Server) CheckFileInFolder(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func (server *Server) CheckFileInServerX(filename string) bool {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	for _, filename2 := range server.Filenames {
		if filename == filename2 {
			return true
		}
	}
	return false
}

func (server *Server) CheckStationIdxInServerX(stationIdx int) bool {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	if _, ok := server.StationIdx2Filename[uint16(stationIdx)]; ok {
		return true
	}
	return false
}
