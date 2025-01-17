# Build Stage
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

# Final Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

# Expose the port for the backend
EXPOSE 8090

# Command to run the Go server
CMD ["./server"]
