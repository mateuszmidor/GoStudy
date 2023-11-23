# syntax=docker/dockerfile:1
FROM golang:1.21 as BUILDER
RUN CGO_ENABLED=0 go install -ldflags="-extldflags=-static" github.com/go-delve/delve/cmd/dlv@latest

FROM alpine
WORKDIR /
COPY --from=BUILDER /go/bin/dlv .
EXPOSE 2345 2345
CMD ["/bin/sh"]