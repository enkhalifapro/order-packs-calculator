FROM golang:alpine AS builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/order-pack-calculator

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/order-pack-calculator main.go

# Start fresh from a smaller image
FROM alpine:latest
RUN apk add ca-certificates

COPY --from=builder /tmp/order-pack-calculator/out/order-pack-calculator /app/order-pack-calculator

# This container exposes port 8090 to the outside world
EXPOSE 8090

CMD ["/app/order-pack-calculator"]