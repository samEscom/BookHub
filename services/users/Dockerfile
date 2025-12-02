# -------------------------
# 1. Builder stage
# -------------------------
FROM golang:1.25-alpine AS builder

# Install git (needed for downloading private deps or modules)
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files first (for caching)
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the service
COPY . .

# Build static binary (no CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -o users-service ./cmd/main.go

# -------------------------
# 2. Final minimal image
# -------------------------
FROM alpine:3.20

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/users-service .

# Runtime config
EXPOSE 8081

CMD ["./users-service"]
