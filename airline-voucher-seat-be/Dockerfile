# Build Stage
FROM golang:1.24.3-alpine AS builder

RUN apk update && apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /app/airline-voucher-seat-be cmd/main/main.go

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/airline-voucher-seat-be .

EXPOSE 8080

ENTRYPOINT ["./airline-voucher-seat-be"]
