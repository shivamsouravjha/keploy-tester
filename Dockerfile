# Use an official Golang runtime as a base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 4000

# Run the application
CMD ["/app/main"]
