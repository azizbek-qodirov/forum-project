# Use a multi-stage build
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o forum .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/forum .
COPY .env .env

# Expose the port on which your application runs
EXPOSE 50051

# Command to run your application
CMD ["./forum"]
