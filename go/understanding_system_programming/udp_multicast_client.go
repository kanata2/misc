package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("server %v\n", remoteAddress)
		fmt.Printf("now    %s\n", string(buffer[:length]))
	}

}
