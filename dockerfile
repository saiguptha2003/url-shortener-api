FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod init url-shortener-api
RUN go mod tidy
RUN go build -o url-shortener-api .

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/url-shortener .
CMD ["./url-shortener"]
