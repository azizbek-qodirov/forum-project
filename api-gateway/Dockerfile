FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway .

FROM alpine:latest

WORKDIR /root/      

COPY --from=builder /app/api-gateway .
#COPY .env .env

EXPOSE 9090

CMD ["./api-gateway"]
