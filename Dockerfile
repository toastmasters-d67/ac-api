FROM golang:alpine AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

# Path: /app/main
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
RUN apk update && apk add tzdata

CMD ["./main"]
