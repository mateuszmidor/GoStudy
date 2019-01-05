# AWS Hello Server
This app is a simple http server designed to be deployable on AWS Beanstalk.
AWS builds the GO application as follows:
    go build -o bin/application application.go
, so there must be a "application.go" main file in the project root dir.