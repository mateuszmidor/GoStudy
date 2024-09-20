package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
	from    = "sender@example.com"
	to      = "recipient@example.com"
	subject = "Welcome to Our Service"
	body    = "Welcome to our service! We're excited to have you on board."
)

func main() {
	addr := flag.String("addr", "localhost:25", "SMTP server address in format host:port")
	useTLS := flag.Bool("tls", false, "Use TLS for email sending")
	flag.Parse()

	host, portStr, err := net.SplitHostPort(*addr)
	if err != nil {
		log.Fatalf("Invalid address format: %v", err)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	server := ServerConfig{
		Host:   host,
		Port:   port,
		UseTLS: *useTLS,
	}
	auth := AuthNone()

	err = SendEmail(context.Background(), from, to, subject, body, server, auth)

	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Printf("Email sent successfully through %s:%d, using tls: %t\n", host, port, *useTLS)
}
