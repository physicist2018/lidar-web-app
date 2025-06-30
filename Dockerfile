# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go

# Final image
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/myapp .
EXPOSE 8080

CMD ["./myapp"]