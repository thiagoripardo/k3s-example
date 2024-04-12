# Stage 1: Build the application
FROM golang:1.22 AS builder

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 go build -ldflags '-s -w' -o /build/app

# Stage 2: Copy the application to a new image
FROM gcr.io/distroless/static-debian12

WORKDIR /
COPY --from=builder /build/app /bin/app

# Stage 3: Set the command to run the application
CMD ["/bin/app"]