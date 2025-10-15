# Multi-stage build for Psiphon Proxy
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o psiphon-proxy .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D -s /bin/sh psiphon

# Create application directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/psiphon-proxy .

# Change ownership to non-root user
RUN chown -R psiphon:psiphon /app
USER psiphon

# Expose default port
EXPOSE 1080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD nc -z localhost 1080 || exit 1

# Run the application
CMD ["./psiphon-proxy"]