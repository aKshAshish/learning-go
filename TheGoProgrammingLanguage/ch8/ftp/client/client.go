package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("Failed to connect; error %v", err)
	}
	defer conn.Close()
	go forward(conn)
	revert(conn)
}

func revert(conn net.Conn) {
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatalf("Error occured while reading response; error %v", err)
	}
}

func forward(conn net.Conn) {
	input := bufio.NewScanner(os.Stdin)
	log.Print("Enter new command:")
	for input.Scan() {
		command := input.Text()
		if _, err := io.WriteString(conn, command+"\n"); err != nil {
			log.Fatalf("Error while sending command %s; %v", command, err)
		}
		log.Print("Enter new command:")
	}
}
