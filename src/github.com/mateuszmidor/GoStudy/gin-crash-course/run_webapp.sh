#!/usr/bin/env bash

function runWebBrowser() {
    sleep 3
    firefox localhost:8080/view/videos
}

go mod download # may take some time
runWebBrowser &
go run .
