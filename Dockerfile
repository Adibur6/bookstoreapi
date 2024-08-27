# Stage 1: Build the application
FROM golang:1.22.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app
RUN ls -l /app
# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
#RUN go build -o ./bookapi .
#Most important line needed for running in minimal base image
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bookapi .

# Stage 2: Create a lightweight image to run the application
FROM alpine:3.18

# Set the Current Working Directory inside the container
WORKDIR /root

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/bookapi .
#RUN ls -l /app
# Add any additional required dependencies (e.g., for bash, you might need busybox)
RUN apk --no-cache add bash

# Expose the port that the app runs on
EXPOSE 8081

# Set up a command to run the application
CMD ["./bookapi", "start", "-p", "8081"]
