package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/protobuf/vault"
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
	validateEndpoint := vault.MakeValidateEndpoints(srv)
	endpoints := vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
	go func() {
		log.Println("http:", *httpAddr)
		handler := vault.NewHTTPServer(ctx, endpoints)
		errCHan <- http.ListenAndServe(*httpAddr, handler)
	}()
	go func() {
		listener, err:= net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- errChan
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler:= vault.NewGRPCServer(ctx, endpoints)
		gRPCServer:= grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()
	log.Fatalln(<-errChan)
}
