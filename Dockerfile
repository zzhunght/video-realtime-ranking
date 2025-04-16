FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp ./cmd

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/config.yaml .

CMD ["./myapp"]
