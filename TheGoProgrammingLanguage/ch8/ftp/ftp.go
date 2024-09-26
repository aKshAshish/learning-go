package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Println("Starting FTP Server...")
	server, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("Failed to start the server; error %v", err)
	}
	defer server.Close()
	log.Println("Listening on port 9000")

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("Connection failed; err %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	log.Printf("Connected to %v", c.RemoteAddr())
	input := bufio.NewScanner(c)
	for input.Scan() {
		command := input.Text()
		go handleCommand(command, c)
	}
	log.Println("Connection Closed to:", c.RemoteAddr())
}

func handleCommand(s string, conn net.Conn) {
	splits := strings.Split(s, " ")
	command := splits[0]
	switch command {
	case "close":
		io.WriteString(conn, "Roger! Closing the connection.\n")
		conn.Close()

	case "cd":
		if len(splits) < 2 {
			io.WriteString(conn, "Provide the folder path")
			return
		}
		if err := os.Chdir(splits[1]); err != nil {
			io.WriteString(conn, fmt.Sprintf("Error occured while changeing directory; error %v\n", err))
		}
	case "ls":
		res, err := handleList()
		if err != nil {
			io.WriteString(conn, fmt.Sprintf("Error occured while changeing directory; error %v\n", err))
			return
		}
		io.WriteString(conn, res)
	default:
		log.Println(command)
	}
}

func handleList() (string, error) {
	result := ""
	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", fmt.Errorf("error occured listing directories")
	}
	dirs, err := os.ReadDir(cwd)
	if err != nil {
		return "", fmt.Errorf("error occured listing directories")
	}
	for _, dir := range dirs {
		result += fmt.Sprintf("%v\n", dir)
	}
	return result, nil
}
