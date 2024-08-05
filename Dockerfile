# Build stage
FROM golang:1.22-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o ascii-art-web .

EXPOSE 8080

CMD ["./ascii-art-web"]