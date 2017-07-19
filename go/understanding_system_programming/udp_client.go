package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("sending to server")
	_, err = conn.Write([]byte("hello world from client"))
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("received: %s\n", string(buffer[:length]))
}
