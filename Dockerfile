# Use the official Golang image as a base for building the application
FROM golang:1.18 AS builder

# Set the GOPATH and create necessary directories
RUN mkdir -p $GOPATH/src/a/b/Yuki-bot 
WORKDIR $GOPATH/src/a/b/Yuki-bot

# Copy the local package files to the container's workspace
COPY . ./

# Install dependencies and build the application
RUN go mod vendor && \
    make build

# Move the built binary to the root directory
RUN mv ./bin/Yuki-bot /

# Expose port 80
EXPOSE 80

# Final stage - use a minimal image for production
FROM alpine

# Copy the compiled binary from the builder stage
COPY --from=builder /Yuki-bot .

# Copy the src folder from the builder stage if needed
COPY --from=builder /go/src/a/b/Yuki-bot/src /src

# Command to run the binary
ENTRYPOINT ["/Yuki-bot"]
