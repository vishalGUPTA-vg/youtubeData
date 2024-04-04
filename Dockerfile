# Build stage
FROM golang:1.21.3-alpine AS builder

WORKDIR /app
COPY . .

# Build the application
RUN go build -o main main.go

# Run stage
FROM alpine:3.13

WORKDIR /app

# Copy the built executable and other necessary files
COPY --from=builder /app/main .
COPY dev.json .

# Expose the port the application runs on
EXPOSE 8080
# Command to run the executable
CMD ["./main"]

