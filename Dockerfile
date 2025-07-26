FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/api/cmd/demo/

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY sqlc/schema /app/schema
EXPOSE 8080
CMD ["./main"]
