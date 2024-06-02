# Use the official Golang image as the base image
FROM golang:latest AS builder

# Set the current working directory inside the container
WORKDIR /server

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image to reduce the final image size
FROM alpine:latest

# Copy the executable from the builder stage to the final stage
COPY --from=builder /server/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]