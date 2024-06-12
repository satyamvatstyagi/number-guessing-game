# Start from the 1.20 golang base image
FROM golang:1.20 AS builder

# Add Maintainer Info
LABEL maintainer="Satyam Vats <satyam.vats@nagarro.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /number-guessing-game ./main.go

# Remove the source code
RUN rm -rf /application/*

# Start a new stage from scratch to build a small image for running the application
FROM alpine:3.11.3

RUN addgroup -g 1000 noroot
RUN adduser -u 1000 -G noroot -h /home/noroot -D noroot
RUN mkdir /home/noroot/app
WORKDIR /home/noroot/app

COPY --from=builder /number-guessing-game /home/noroot/app

# Run the executable
CMD ["./number-guessing-game"]
