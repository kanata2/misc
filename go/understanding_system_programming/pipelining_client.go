package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn
	var err error
	requests := make(chan *http.Request, len(sendMessages))
	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access: %d\n", current)
	defer conn.Close()
	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		if lastMessage {
		}
	}
}
