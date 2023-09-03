package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// to tell the client-side what host to connect to
type indexPageData struct {
	Host string
}

// to send the client-side the current time
type timeData struct {
	Time string `json:"time"`
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}
var indexTemplate *template.Template = template.Must(template.ParseFiles("index.html"))

// serve HTML
func handleIndex(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, indexPageData{Host: r.Host})
}

// serve WebSocket
func handleTime(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("WebSocket error:", err)
		return
	}

	// serve time over WebSocket in 1 second intervals
	fmt.Println("New client connected")
	go func() {
		for {
			currTime := time.Now().Format("15:04:05")
			data := timeData{Time: currTime}
			err := socket.WriteJSON(data)
			if err != nil {
				fmt.Println("Client disconnected")
				break
			}
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/time", handleTime)

	fmt.Println("awaiting clients at localhost:9000...")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
