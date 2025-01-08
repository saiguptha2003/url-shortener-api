FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod init url-shortener-api
RUN go mod tidy
RUN go build -o url-shortener-api .

