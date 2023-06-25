# Start from a minimal base image with Go installed
FROM golang:1.18.6-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o server .

# Start from a new minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=build /app/server .

# Expose the port on which the server listens
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./server"]
