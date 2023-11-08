# Use the official Golang image as a base image for building
FROM golang:alpine AS builder

# Install git
RUN apk update && apk add --no-cache git=~2

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application
COPY . .
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# Start a new stage to create a minimal image
FROM alpine:3

# Set the working directory
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/main .
COPY .env /app
# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]

