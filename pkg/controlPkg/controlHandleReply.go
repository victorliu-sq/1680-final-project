package controlPkg

import (
	"fmt"
	"log"
	"os"
	"snowcast/pkg/protocol/reply"
)

func (cc *ClientControl) HandleWelComeReply(bytes []byte) {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	wMsg := reply.UnmarsharlWelcomeMsg(bytes)
	// fmt.Printf("Welcome to Snowcast! The server has %d stations.\n", wMsg.GetNumStations())
	// 1. check whether the welcome is received before an Hello Msg
	if !cc.IsHelloSent {
		// log.Fatalln("The Welcome Msg is received before an Hello Msg")
		os.Exit(0)
	}
	// 2. check whether multiple welcomes
	if cc.IsWelcomeRev {
		// log.Fatalln("Multiple Welcome Msgs are received")
		os.Exit(0)
	}
	cc.IsWelcomeRev = true
	fmt.Printf("Welcome to Snowcast! The server has %v stations\n", wMsg.GetNumStations())
}

func (cc *ClientControl) HandleAnnounceReply(bytes []byte) {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	if !cc.IsSetStationSent {
		// log.Fatalln("The announce reply is received before an setStation Msg")
		os.Exit(0)
	}
	// log.Println("Receive an AnnounceMsg")
	size := uint8(bytes[1])
	// log.Printf("length of string is %v\n", size)
	songnameBytes := bytes[2 : 2+size]
	fmt.Printf("New song announced: %v\n", string(songnameBytes))
	// check whether announce reply is received before an setStation Msg
}

func (cc *ClientControl) HandleInvalidReply(bytes []byte) {
	log.Println("Receive an Invalid Reply Msg")
	size := uint8(bytes[1])
	// log.Printf("length of string is %v\n", size)
	replyStringBytes := bytes[2 : 2+size]
	fmt.Printf("Receive an Invalid Reply Msg, invalid reason is %v\n", string(replyStringBytes))
	// End this process of client
	os.Exit(0)
}

func (cc *ClientControl) HandleNewStationReply(bytes []byte) {
	nMsg := reply.UnmarsharlNewStationMsg(bytes)
	fmt.Printf("New Station! The server now has %d stations.\n", nMsg.StationNum)
}

func (cc *ClientControl) HandleStationShutDownReply(bytes []byte) {
	sMsg := reply.UnmarsharlStationShutDownMsg(bytes)
	fmt.Printf("StationIdx %v has been shutdown :( Try another one.\n", sMsg.StationIdx)
}
