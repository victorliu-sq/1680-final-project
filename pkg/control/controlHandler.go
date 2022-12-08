package control

import (
	"fmt"
	"os"
	"snowcast/pkg/protocol"
)

func (cc *ClientControl) HandleClientMsg() {
	for {
		clientCLI := <-cc.ClientMsgChan
		switch clientCLI.MsgType {
		case protocol.TypeSetStation:
			fmt.Println("hello")
			// clientControl.SendSetStationMsg(clientCLI.StationNum)
		case protocol.TypeQuitClient:
			os.Exit(0)
		}
	}
}
