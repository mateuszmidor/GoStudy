package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/mateuszmidor/GoStudy/connectgo/pingpong"
	"github.com/mateuszmidor/GoStudy/connectgo/pingpong/pingpongconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const NET_ADDR = ":9000"

// Server implements pingpongconnect.PingPongHandler interface generated from .proto
type Server struct {
	pingpongconnect.UnimplementedPingPongHandler // returns error from all methods
}

// RpcPing implements pingpongconnect.PingPongHandler interface
func (s *Server) RpcPing(_ context.Context, r *connect.Request[pingpong.Message]) (*connect.Response[pingpong.Message], error) {
	log.Printf("server received request: %s", r.Msg.GetBody())
	msg := pingpong.Message{Body: "Pong!"}
	res := connect.NewResponse(&msg)
	return res, nil
}

func main() {
	go server()
	time.Sleep(time.Second) // lousy! but works.
	client()

	time.Sleep(time.Hour) // keep server alive for playing around with curl
}

func server() {
	mux := http.NewServeMux()
	mux.Handle(pingpongconnect.NewPingPongHandler(&Server{}))
	err := http.ListenAndServe(NET_ADDR, h2c.NewHandler(mux, &http2.Server{}))
	log.Fatalf("listen failed: %v", err)
}

func client() {
	client := pingpongconnect.NewPingPongClient(http.DefaultClient, "http://"+NET_ADDR)
	msg := pingpong.Message{Body: "Ping!"}
	req := connect.NewRequest(&msg)
	res, err := client.RpcPing(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("client received response: %s", res.Msg.GetBody())
}
