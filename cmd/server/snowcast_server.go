package main

import (
	"log"
	"os"
	"snowcast/pkg/server"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Usage %v <server listen port> + <filename> ...", os.Args[0])
	}
	server := server.Server{}
	server.Make(os.Args)
}
