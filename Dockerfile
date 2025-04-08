FROM golang:1.22-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .

COPY --from=builder /app/config/.env ./config/.env
ENV ENV=production
EXPOSE 8080

CMD ["./app"]
