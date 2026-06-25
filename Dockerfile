# syntax=docker/dockerfile:1.7

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/go-ftp-server-linux-amd64 .

# Windows binary
RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
    go build -o /out/go-ftp-server-windows-amd64.exe .


# FINAL STAGE (export only)
FROM scratch AS export
COPY --from=builder /out /out
