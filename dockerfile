# Stage 1: Build the application
FROM golang:1.20 AS builder


WORKDIR /app
COPY . /app/

# Copy only the necessary files to leverage Docker's caching
RUN go mod init url-shortener-api
RUN go mod tidy



# Build the application
RUN go build -o url-shortener-api .

# Stage 2: Create a minimal container with the built binary
FROM gcr.io/distroless/base-debian10:latest

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener-api /url-shortener-api

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/url-shortener-api"]
