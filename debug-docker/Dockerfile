# syntax=docker/dockerfile:1
FROM golang:1.21

WORKDIR /app
COPY . .

# install dlv at: /go/bin/dlv
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# build the app itself, without inlining and other optimizations that might hinder debugging
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o HelloWorld

# expose dlv port
EXPOSE 2345

# by default, just run the app
CMD ["/app/HelloWorld"]