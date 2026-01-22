FROM golang:1.22-alpine AS builder

WORKDIR /app

# 1. Copy everything
COPY . .

# 2. DEBUG: List all files recursively so we can see the structure in the logs
RUN ls -R

# 3. Check for go.mod explicitly
RUN if [ ! -f go.mod ]; then echo "ERROR: go.mod not found in /app"; exit 1; fi

RUN go mod download
RUN go build -o main ./cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
