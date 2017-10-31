package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Broker struct {
	Notifier       chan []byte
	newClients     chan chan []byte
	closingClients chan chan []byte
	clients        map[chan []byte]bool
}

func NewServer() *Broker {
	broker := &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}
	go broker.listen()
	return broker
}

func (b *Broker) listen() {
	for {
		select {
		case s := <-b.newClients:
			b.clients[s] = true
			log.Printf("Added client. %d registered", len(b.clients))
		case s := <-b.closingClients:
			delete(b.clients, s)
			log.Printf("Removed client. %d registered", len(b.clients))
		case event := <-b.Notifier:
			for clientMsgChan := range b.clients {
				clientMsgChan <- event
			}
		}
	}
}
func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	msgChan := make(chan []byte)
	b.newClients <- msgChan
	defer func() {
		b.closingClients <- msgChan
	}()
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		b.closingClients <- msgChan
	}()
	for {
		fmt.Fprintf(w, "data: %s\n\n", <-msgChan)
		flusher.Flush()
	}
}

func main() {
	broker := NewServer()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			eventStr := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Receiving event")
			broker.Notifier <- []byte(eventStr)
		}
	}()
	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", broker))
}
