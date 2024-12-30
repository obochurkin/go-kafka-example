# Dockerfile
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the server binary
ARG TARGET=server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/"$TARGET"

# Create a smaller image for the server/consumer
FROM alpine:3.18

COPY --from=builder /bin/app /bin/app

ENTRYPOINT ["/bin/app"]
