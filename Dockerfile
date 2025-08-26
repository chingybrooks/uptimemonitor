FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o uptimemonitor \
    -ldflags "-X uptimemonitor/pkg/version.Version=${VERSION:-$(git describe --tags)}" \
    ./cmd/uptimemonitor

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/uptimemonitor .

EXPOSE 3000

CMD ["./uptimemonitor","-addr=:$PORT"]
