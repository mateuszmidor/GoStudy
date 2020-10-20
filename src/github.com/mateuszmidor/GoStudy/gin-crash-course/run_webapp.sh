#!/usr/bin/env bash

function runWebBrowser() {
    sleep 3
    firefox localhost:8080/view/videos
}

runWebBrowser &
go run .