package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func writeToConn(sessionResps chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for sessionResp := range sessionResps {
		resp := <-sessionResp
		resp.Write(conn)
		close(sessionResp)
	}
}

func handleRequest(req *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	content := "Hello, World\n"
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}
	resultReceiver <- response
}

func porcessSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)
	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse
		go handleRequest(request, sessionResponse)
	}
}
