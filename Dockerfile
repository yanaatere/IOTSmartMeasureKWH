# Start with the official Go image for building the application
FROM golang:1.22.1 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory from your local machine to the container
COPY . .

# Build the Go application (adjust the output path if needed)
RUN go build -o main .

# Create a minimal runtime image using Alpine
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage to the runtime stage
COPY --from=builder /app/main .

# Run the compiled Go binary
CMD ["./main"]
