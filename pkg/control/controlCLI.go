package control

import (
	"bufio"
	"fmt"
	"os"
	"snowcast/pkg/protocol"
	"strconv"
)

func (cc *ClientControl) ScanClientCLI() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if num, err := strconv.Atoi(line); err == nil {
				fmt.Println("Try to set station", line)
				clientCLI := protocol.NewClientMsg(protocol.TypeSetStation, uint16(num))
				cc.ClientMsgChan <- clientCLI
			} else if len(line) == 1 && line[0] == 'q' {
				clientCLI := protocol.NewClientMsg(protocol.TypeQuitClient, uint16(0))
				cc.ClientMsgChan <- clientCLI
			}
		}
	}
}

func (cc *ClientControl) HandleClientMsg() {
	for {
		clientCLI := <-cc.ClientMsgChan
		switch clientCLI.MsgType {
		case protocol.TypeSetStation:
			fmt.Println("hello, try to set station")
			cc.CallSetStationRPC(clientCLI.StationNum)
		case protocol.TypeQuitClient:
			os.Exit(0)
		}
	}
}
