# Build stage
FROM golang:1.23.5-alpine AS builder

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN go build -o server ./cmd/server

# Runtime stage  
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run
CMD ["./server"] 