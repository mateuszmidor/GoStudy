package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn // 2 way connection server <-> browser
	send   chan []byte     // data to be sent to browser
	room   *room           // room this client is connected to
}

// Receive messages coming from browser
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

// Send messages stored in output channel to the browser
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
