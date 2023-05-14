# Use an official Golang image as the base image
FROM golang:1.20-alpine

# Set the DB_HOST environment variable
ENV DB_HOST="barnyard-database1"

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

RUN go mod tidy
# Download Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o main .



# Expose the port that the API will be served on
EXPOSE 5000

# Start the API
CMD ["./main"]