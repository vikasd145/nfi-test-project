# Start from a minimal base image with Go installed
FROM golang:1.18.6-alpine AS build

# Install build tools
RUN apk add --no-cache build-base

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Run the test cases
RUN go test -v ./...

RUN mkdir -p bin

# Build the Go application if tests pass
RUN go build -o bin ./...

# Start from a new minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=build /app/bin .
COPY --from=build /app/configs/config.yaml config.yaml

# Expose the port on which the server listens
EXPOSE 8080

# Set the entry point for the container
CMD ["./nfi-test-project"]
