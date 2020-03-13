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
	"time"

	ratelimitkit "github.com/go-kit/kit/ratelimit"
	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/r10-protobuf/vault"
	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/r10-protobuf/vault/pb"
	"golang.org/x/time/rate"
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

	limit := rate.NewLimiter(rate.Every(time.Second), 5) // allow 5 requests/sec
	hashEndpoint := vault.MakeHashEndpoint(srv)
	hashEndpoint = ratelimitkit.NewErroringLimiter(limit)(hashEndpoint)
	validateEndpoint := vault.MakeValidateEndpoint(srv)
	validateEndpoint = ratelimitkit.NewErroringLimiter(limit)(validateEndpoint)
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
