# Use an official Golang runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Install Git
RUN apt-get update && apt-get install -y git

# Install Redis
RUN apt-get install -y redis-server

# Install MongoDB
RUN apt-get install -y mongodb

# Install necessary Go packages for Gin and MongoDB
RUN go get -u github.com/gin-gonic/gin
RUN go get go.mongodb.org/mongo-driver/mongo

# Build the Go app
RUN go mod download
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

