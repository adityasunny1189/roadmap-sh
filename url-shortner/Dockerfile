# Use the official Golang image to build the app
FROM golang:1.20-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies (this will be cached if go.mod/go.sum do not change)
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o url-shortener ./main.go

# Use a smaller base image for the final output
FROM alpine:latest

# Set environment variables (can be overridden at runtime using docker-compose or docker run)
ENV DB_USER=your_db_user
ENV DB_PASSWORD=your_db_password

# Set the working directory in the container
WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=build /app/url-shortener .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./url-shortener"]
