package main

import "os"

// GetPort returns port number from env variable "PORT" if exists, 8080 otherwise
func GetPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return port
	}
	return "8080"
}

// GetAPIEndpoint returns api endpoint from env variable "API_ENDPOINT" if exists, http://localhost:8080 otherwise
func GetAPIEndpoint() string {
	endpoint := os.Getenv("API_ENDPOINT")
	if endpoint != "" {
		return endpoint
	}
	return "http://localhost:" + GetPort()
}
