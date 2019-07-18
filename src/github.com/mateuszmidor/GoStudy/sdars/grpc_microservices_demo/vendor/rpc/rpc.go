package rpc

import (
	fmt "fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func LogCallResult(err error) {
	if err != nil {
		// fmt.Printf("Call failed with %v\n", err)
		fmt.Println("call failed")
	}
}

func ConnectGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %s %v\n", addr, err)
	} else {
		log.Println("connection ready.")
	}
	return conn
}

func MakeGrpcServer(addr string) (net.Listener, *grpc.Server) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %s %v\n", addr, err)
	}
	log.Printf("server listening at %s\n", addr)
	grpcServer := grpc.NewServer()
	return lis, grpcServer
}
