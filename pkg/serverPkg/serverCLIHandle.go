package serverPkg

import (
	"fmt"
	"log"
	"os"
	"snowcast/pkg/protocol/CLI"
)

func (server *Server) HandleCLIPrintClients() {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	for idx, filename := range server.Filenames {
		output := fmt.Sprintf("%v,%v", idx, filename)
		for controlAddr, _ := range server.StationIdx2Controls[uint16(idx)] {
			output += ","
			output += server.Control2Listener[controlAddr]
		}
		fmt.Println(output)
	}
}

func (server *Server) HandleCLIPrintToFile(serverCLI CLI.ServerCLIMsg) {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	// create a file
	filename := serverCLI.Filename
	// filePath := fmt.Sprintf("./%v", filename)
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	// for a line for each file
	for idx, filename := range server.Filenames {
		output := fmt.Sprintf("%v,%v", idx, filename)
		for controlAddr, _ := range server.StationIdx2Controls[uint16(idx)] {
			output += ","
			output += server.Control2Listener[controlAddr]
		}
		output += "\n"
		// write into file
		_, err = f.WriteString(output)
		if err != nil {
			log.Println(err)
		}
	}
}

func (server *Server) HandleCLIQuitServer() {
	fmt.Println("server will quit")
	os.Exit(0)
}

/*
	To add a file, all variable we need to deal with are:
	Filenames
	Filename2Chunks
	StationIdx2Filename
	StationIdx2Controls
*/

func (server *Server) HandleCLIAddFile(serverCLI CLI.ServerCLIMsg) {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	filename := serverCLI.Filename
	stationIdx := uint16(len(server.Filenames))
	// chunks := server.ToChunks(filename)
	server.Filenames = append(server.Filenames, filename)
	// server.Filename2Chunks[filename] = chunks
	server.StationIdx2Filename[stationIdx] = filename
	server.StationIdx2Controls[stationIdx] = map[string]int{}
	// Start a DaemonStation for this stationIdx
	go server.DaemonStation(int(stationIdx), filename)
	fmt.Printf("Add one station for file %v successfully\n", serverCLI.Filename)
	server.BroadcastNewStationMsg(uint16(len(server.Filenames)))
}

/*
	To add a file, all variable we need to deal with are:
	Filenames
	Filename2Chunks
	StationIdx2Filename
	StationIdx2Controls
*/

func (server *Server) HandleCLIRemoveStationIdx(serverCLI CLI.ServerCLIMsg) {
	server.Mu.Lock()
	defer server.Mu.Unlock()
	stationIdxToDel := serverCLI.StationNum
	// filenameToDel := server.Filenames[stationIdxToDel]

	server.BroadcasStationShutDownMsg(stationIdxToDel)
	// filename2chunks
	// delete(server.Filename2Chunks, filenameToDel)
	// StationIdx2Controls
	delete(server.StationIdx2Controls, stationIdxToDel)
	// StationIdx2Filename, remove this will stop the corresponding Daemon to send UDP data
	delete(server.StationIdx2Filename, stationIdxToDel)
	fmt.Printf("Remove one stationIdx %v successfully\n", serverCLI.StationNum)

	// filenames: if oldStationIdx > stationIdxToDel, newStationIdx = old - 1
	// fmt.Println("filenames before deletion", server.Filenames)
	for oldStationIdx, filename := range server.Filenames {
		if oldStationIdx <= int(stationIdxToDel) {
			continue
		}
		// Filenames
		newStationIdx := oldStationIdx - 1
		server.Filenames[newStationIdx] = filename
		// StationIdx2Controls
		controls := server.StationIdx2Controls[uint16(oldStationIdx)]
		delete(server.StationIdx2Controls, uint16(oldStationIdx))
		server.StationIdx2Controls[uint16(newStationIdx)] = controls
		// StationIdx2Filename
		filename := server.StationIdx2Filename[uint16(oldStationIdx)]
		delete(server.StationIdx2Filename, uint16(oldStationIdx))
		server.StationIdx2Filename[uint16(newStationIdx)] = filename
	}
	server.Filenames = server.Filenames[:len(server.Filenames)-1]
	// fmt.Println("filenames after deletion", server.Filenames)
}
