FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy only the module files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest of the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/api/main.go

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
