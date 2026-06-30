FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-s -w -buildid=" \
    -o /out/go-ftp-server-linux-amd64 .

RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-s -w -buildid=" \
    -o /out/go-ftp-server-windows-amd64.exe .

FROM scratch AS export
COPY --from=builder /out /out
