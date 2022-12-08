package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const chunkSize = 4096

func main() {
	// check os.Args #
	if len(os.Args) < 2 {
		log.Fatalf("Usage:  %s <port number>", os.Args[0])
	}
	// 1. Specify port number and get UDP addr
	addr, err := net.ResolveUDPAddr("udp", ToColonPortNumber(os.Args[1]))
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(addr)
	// 2. Listen on the udp port and get 1 udp conn
	conn, err := net.ListenUDP("udp", addr)
	// conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	// 3. Receive data from server
	// log.Printf("Connect with server %v\n", conn.RemoteAddr().String())
	for {
		bytes := make([]byte, chunkSize)
		n, err := conn.Read(bytes)
		if err != nil {
			log.Fatalln(err)
		}
		bytes = bytes[:n]
		// fmt.Println(bytes)
		os.Stdout.Write(bytes)
	}
}

func ToColonPortNumber(portNum string) string {
	return fmt.Sprintf(":%v", portNum)
}
