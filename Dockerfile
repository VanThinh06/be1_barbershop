# Build stage
FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE  8080
CMD [ "/app/main" ]