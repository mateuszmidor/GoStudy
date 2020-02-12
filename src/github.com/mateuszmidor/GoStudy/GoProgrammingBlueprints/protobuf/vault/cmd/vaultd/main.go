package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/protobuf/vault"
	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/protobuf/vault/pb"
	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "HTTP port nubmer")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC port number")
	)
	flag.Parse()
	ctx := context.Background()
	srv := vault.NewService()
	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	hashEndpoint := vault.MakeHashEndpoint(srv)
	validateEndpoint := vault.MakeValidateEndpoint(srv)
	endpoints := vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
	go func() {
		log.Println("http:", *httpAddr)
		handler := vault.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := vault.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()
	log.Fatalln(<-errChan)
}
