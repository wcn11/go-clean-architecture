# Builder stage
FROM golang:1.20.7-alpine3.17 as builder

# Install dependencies
RUN apk update && apk add --no-cache git make bash build-base

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the application
RUN make build

# Distribution stage
FROM alpine:latest

# Install necessary packages
RUN apk update && apk add --no-cache tzdata

# Create the working directory
WORKDIR /app 

# Expose the application port
EXPOSE 9090

# Copy the built application from the builder stage
COPY --from=builder /app/engine /app/

# Command to run the application
CMD ["/app/engine"]
