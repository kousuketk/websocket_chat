FROM golang:latest
WORKDIR /app

COPY app/go.mod .
COPY app/go.sum .

RUN go mod download