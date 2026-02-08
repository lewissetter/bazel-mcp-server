# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bazel-mcp-server ./cmd/bazel-mcp-server

# Final stage
FROM gcr.io/bazel-public/bazel:latest

# Copy the binary from builder
COPY --from=builder /build/bazel-mcp-server /usr/local/bin/bazel-mcp-server

# Set up a working directory
WORKDIR /workspace

# Run as non-root user
USER bazel

# Run the server
ENTRYPOINT ["bazel-mcp-server"]
