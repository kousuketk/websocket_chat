FROM golang:latest
WORKDIR /app

COPY server/go.mod .
COPY server/go.sum .

RUN go mod download