

# Build stage
FROM golang:1.20.12-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

# run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/api .
COPY app.env .

EXPOSE 8081
CMD ["/app/api"]