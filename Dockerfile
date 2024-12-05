# Start from the official Golang image
FROM golang:1.21.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose port 8081 for the API
EXPOSE 8081

# Run the binary as the entrypoint
CMD ["./main"]
