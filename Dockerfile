# Use official Go image as a base
FROM golang:1.21-alpine

# Create working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port your app runs on
EXPOSE 8080

# Run the executable
CMD [ "./main" ]
