package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := os.Args[1]
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // Blocks here until a connection is received
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Printf("Client disconnected, %v \n", c.LocalAddr())
			return
		}
		time.Sleep(1 * time.Second)
	}
}
