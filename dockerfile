FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod init url-shortener
RUN go mod tidy
RUN go build -o url-shortener .

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/url-shortener .
CMD ["./url-shortener"]
