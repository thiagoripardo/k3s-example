# Use a multi-stage build
FROM golang:1.22 AS builder

WORKDIR /build
COPY . .

# Build for amd64
RUN CGO_ENABLED=0 GOARCH=amd64 go build -ldflags '-s -w' -o /build/app-amd64

# Build for arm64
RUN CGO_ENABLED=0 GOARCH=arm64 go build -ldflags '-s -w' -o /build/app-arm64

# Final stage
FROM gcr.io/distroless/static-debian12

# Copy the binaries from the previous stage
COPY --from=builder /build/app-amd64 /bin/app-amd64
COPY --from=builder /build/app-arm64 /bin/app-arm64

# The CMD instruction is omitted as it depends on the deployment environment
