# Start from the latest golang base image 1
FROM golang:1.19.5-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN cd app/extl/ \
    && ls \
    && pwd \
    && go build -tags musl main.go

# Expose port 80 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app/main/v1/main"]
# CMD ["sh", "-c", "/root/app/extl/main"]