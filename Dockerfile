# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application on port 8080
EXPOSE 8080

# Start the application
CMD ["./main"]
