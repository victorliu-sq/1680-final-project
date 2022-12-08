package controlPkg

import (
	"fmt"
	"log"
	"os"
	"snowcast/pkg/protocol/CLI"
	"snowcast/pkg/protocol/reply"
	"time"
)

func (cc *ClientControl) RevReplyMsg() {
	// Set Deadline for Welcome Msg within 100ms
	waitForWelcome := true
	for {
		if waitForWelcome {
			cc.ServerConn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			waitForWelcome = false
		} else {
			cc.ServerConn.SetReadDeadline(time.Now().Add(3000 * time.Second))
		}
		// Receive ReplyMsg from server
		bytes := make([]byte, 128)
		byteNum, err := cc.ServerConn.Read(bytes)
		if byteNum < 1 || err != nil {
			log.Fatalln("The server has been closed")
		}
		bytes = bytes[:byteNum]
		// The client have to receive subsequent bytes within 100 milliseconds
		cur := byteNum
		var total int
		if uint8(bytes[0]) == reply.TypeWelcome || uint8(bytes[0]) == reply.TypeNewStation || uint8(bytes[0]) == reply.TypeStationShutDown {
			total = 3
		} else if cur >= 2 && uint8(bytes[0]) == reply.TypeAnnounce || uint8(bytes[0]) == reply.TypeInvalid {
			total = 2 + int(uint8(bytes[1]))
		} else {
			// if we only get announce/invalid type but not size of msg
			// set total to a large num to reset size later
			total = 128
		}
		for cur < total {
			cc.ServerConn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			temp := make([]byte, 128)
			byteNum, err = cc.ServerConn.Read(temp)
			if err != nil {
				log.Fatalln("Problem when trying to read subsequence bytes")
			}
			temp = temp[:byteNum]
			bytes = append(bytes, temp...)
			cur += byteNum
			// update total
			if cur >= 2 && uint8(bytes[0]) == reply.TypeAnnounce || uint8(bytes[0]) == reply.TypeInvalid {
				total = 2 + int(uint8(bytes[1]))
			}
			// fmt.Println(cur)
		}
		replyType := uint8(bytes[0])
		fmt.Printf("Read %v bytes, replyType is %v\n", byteNum, replyType)
		switch replyType {
		case reply.TypeWelcome:
			cc.HandleWelComeReply(CLI.CopyBytes(bytes))
		case reply.TypeAnnounce:
			cc.HandleAnnounceReply(CLI.CopyBytes(bytes))
		case reply.TypeInvalid:
			cc.HandleInvalidReply(CLI.CopyBytes(bytes))
		// Test Case will fail if we use reply.TypeNewStation and TypeStationShutDown
		// ****************************************************************
		// case reply.TypeNewStation:
		// 	cc.HandleNewStationReply(CLI.CopyBytes(bytes))
		// case reply.TypeStationShutDown:
		// 	cc.HandleStationShutDownReply(CLI.CopyBytes(bytes))
		// ****************************************************************
		default:
			os.Exit(0)
		}
	}
}
