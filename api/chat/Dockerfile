# Stage 1: Build the application
FROM golang:1.22 AS builder

WORKDIR /build
COPY . .

RUN make build

# Stage 2: Copy the application to a new image
FROM gcr.io/distroless/static-debian12

WORKDIR /bin
COPY --from=builder /build/bin/chat .

# Stage 3: Set the command to run the application
CMD ["/bin/chat"]
