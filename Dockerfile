FROM golang:1.22.0 AS builder 

WORKDIR /app 

COPY go.mod .
COPY main.go . 

RUN go get 

RUN go build -o bin .

ENTRYPOINT ["/app/bin"] 