FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8090

CMD ["./server"]
