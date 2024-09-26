package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for _, addr := range os.Args[2:] {
		go getTime(addr)
	}

	addr := os.Args[1]
	getTime(addr)
}

func getTime(addr string) {
	c, err := net.Dial("tcp", "localhost:"+addr)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	defer c.Close()
	if _, err := io.Copy(os.Stdout, c); err != nil {
		log.Fatalf("%v", err)
	}
}
