package controlPkg

import (
	"log"
	"snowcast/pkg/protocol/command"
)

func (cc *ClientControl) SendHelloMsg(udpPort string) {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	helloMsg := command.NewHelloMsg(udpPort)
	bytes := helloMsg.Marsharl()
	cc.ServerConn.Write(bytes)
	cc.IsHelloSent = true
}

func (cc *ClientControl) SendSetStationMsg(stationNum uint16) {
	cc.Mu.Lock()
	defer cc.Mu.Unlock()
	setStationMsg := command.NewSetStationMsg(stationNum)
	bytes := setStationMsg.Marsharl()
	// cc.ServerConn.Write(bytes)
	// cc.ServerConn.Write(bytes[:1])
	// time.Sleep(50 * time.Millisecond)
	// cc.ServerConn.Write(bytes[1:])
	cur, total := 0, len(bytes)
	for cur < total {
		n, err := cc.ServerConn.Write(bytes[cur:])
		if err != nil {
			log.Fatalln(err)
		}
		cur += n
	}
	cc.IsSetStationSent = true
	// fmt.Println("Send SetStationMsg successfully!", setStationMsg.StationNumber)
}

func (cc *ClientControl) SendUnknownMsg(portNum string) {
	unknownMsg := command.NewUnknownMsg(portNum)
	bytes := unknownMsg.Marsharl()
	cc.ServerConn.Write(bytes)
}
