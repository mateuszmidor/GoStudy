package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/mateuszmidor/GoStudy/grpc/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const NET_ADDR = ":9000"

// Server implements pingpong.PingPongServer interface generated from .proto
type Server struct {
	pingpong.UnimplementedPingPongServer
}

// RpcPing implements pingpong.PingPongServer interface
func (s *Server) RpcPing(ctx context.Context, msg *pingpong.Message) (*pingpong.Message, error) {
	log.Printf("server received request: %s\n", msg.GetBody())
	return &pingpong.Message{Body: "Pong!"}, nil
}

func main() {
	go server()
	time.Sleep(time.Second) // lousy! but works.
	client()
}

func server() {
	// create TCP listener
	lis, err := net.Listen("tcp", NET_ADDR)
	if err != nil {
		log.Fatal("server listen failed: ", err)
	}

	// create gRPC server waiting for requests
	grpcServer := grpc.NewServer()
	s := Server{}
	pingpong.RegisterPingPongServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("server serve failed: ", err)
	}
}

func client() {
	// dial gRPC server
	conn, err := grpc.Dial(NET_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client dial failed: ", err)
	}
	defer conn.Close()

	// send request to gRPC server
	client := pingpong.NewPingPongClient(conn)
	resp, err := client.RpcPing(context.Background(), &pingpong.Message{Body: "Ping!"})
	if err != nil {
		log.Fatal("client call to server failed: ", err)
	}
	log.Println("client received response:", resp.GetBody())
}
