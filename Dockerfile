# Use the official Golang image as the base image for building
FROM golang:1.23.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app from cmd/main.go
RUN GOOS=linux GOARCH=amd64 go build -o telegram-bot cmd/main.go

# Start a new stage from scratch to keep the image small
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/telegram-bot .

# Ensure it's executable (optional)
RUN chmod +x telegram-bot

# Expose port (if your bot listens on a specific port)
EXPOSE 8080

# Command to run the executable
CMD ["./telegram-bot"]
