FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/cmd/main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/main .

COPY .env .env

CMD ["./main"]