package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn // 2 way connection server <-> browser
	send     chan *message   // data to be sent to browser
	room     *room           // room this client is connected to
	userData map[string]interface{}
}

// Receive messages coming from browser
func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		if avatarURL, ok := c.userData["avatar_url"]; ok {
			msg.AvatarURL = avatarURL.(string)
		}
		c.room.forward <- msg
	}
}

// Send messages stored in output channel to the browser
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
