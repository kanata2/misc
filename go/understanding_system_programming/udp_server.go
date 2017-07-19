package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("hello, world from server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
