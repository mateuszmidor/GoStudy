# syntax=docker/dockerfile:1
FROM golang:alpine3.18

# build the app binary
WORKDIR /app
COPY main.go go.mod go.sum .
RUN go build -o app .

# setup non-root user
ARG USER=regular-user
RUN adduser -D ${USER}
USER ${USER}
ENV PATH=/home/${USER}/.local/bin:$PATH

CMD ["sh", "-c", "id; ./app"]