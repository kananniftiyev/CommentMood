# Use the official Go image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Expose the port that the application will run on
EXPOSE 8080

# Run the application
CMD ["./main"]
