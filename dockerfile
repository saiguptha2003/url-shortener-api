FROM golang:1.15.7-buster

WORKDIR /app
COPY . .
ENTRYPOINT ["go", "run", "main.go"]