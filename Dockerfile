# Build stage - This stage compiles our Go application
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker cache
# This is a best practice - if dependencies don't change, this layer will be cached
COPY go.mod go.sum ./
# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
# CGO_ENABLED=0: Disables CGO (C Go) which creates a static binary
# GOOS=linux: Specifies the target operating system
# -o main: Names the output binary as 'main'
# .: Builds the package in the current directory
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage - This stage creates the minimal runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
CMD ["./main"]
