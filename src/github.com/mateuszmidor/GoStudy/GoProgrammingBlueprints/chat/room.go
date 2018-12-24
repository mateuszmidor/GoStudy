package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/trace"
)

type room struct {
	forward chan []byte

	join chan *client

	leave chan *client

	clients map[*client]bool

	tracer trace.Tracer
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("Do pokoju dołączył nowy klient")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Klient opuścił pokój")
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- wysłano do klienta")
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
