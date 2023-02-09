# AWS Beanstalk Hello Server

This app is a simple http server designed to be deployable on AWS Beanstalk.  
It dumps environment variables to stdout and renders welcome web page.

## Build

AWS builds the GO application as follows:
```sh
go build -o bin/application application.go
```
, so there must be an "application.go" main file in the project root dir.  

## Run

AWS runs the app with exposed port 5000, so it must listen on port 5000.