#!/usr/bin/env bash


trap killserver SIGINT

function installGoPackages() {
    go get github.com/go-kit/kit
    go get golang.org/x/time/rate
    go get -u -v golang.org/x/crypto/bcrypt
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
}

function runserver() {
    go run vault/cmd/vaultd/main.go 2>/dev/null &
    while [[ `pgrep -f vaultd` == "" ]]; do sleep 1; done
    sleep 1
}

function killserver() {
    pkill -f "test|vaultd|go-build" 
}

function runHTTPRequest() {
    echo
    echo "Hashing maciek (HTTP)"
    curl -X POST -s -d '{"password":"maciek"}' http://localhost:8080/hash 
}

function runGRPCRequest() {
    echo
    echo "Hashing maciek (GRPC)"
    go run vault/cmd/vaultcli/main.go hash maciek
}

function runRateLimitCheck() {
    echo
    echo "Rate limit check (max 5/sec)"
    for i in {1..10}; do
        curl -X POST -s -d '{"password":"maciek"}' http://localhost:8080/hash 
        sleep 0.2
        echo
    done
}

#installGoPackages
runserver
runHTTPRequest
runGRPCRequest
runRateLimitCheck
killserver