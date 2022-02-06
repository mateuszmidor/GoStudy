FROM golang as builder
WORKDIR /builder
COPY ./src/ ./
RUN cd ./cmd/finder_web/ && CGO_ENABLED=0 go build -o server .

FROM alpine
WORKDIR /opt/
COPY ./data/*.gz ./csv/
COPY src/cmd/finder_web/data/ data/
EXPOSE 80
COPY --from=builder /builder/cmd/finder_web/server .
ENTRYPOINT ["/opt/server", "-port=80", "-data=./csv"]