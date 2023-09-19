FROM golang:1.21.1-bookworm as builder

WORKDIR /app/rinha-go

COPY . .

RUN GOOS=linux go build -o rinha cmd/main.go

FROM debian:bookworm-slim

COPY --from=builder /app/rinha-go/rinha /usr/local/bin/rinha

ENTRYPOINT [ "rinha", "/var/rinha/source.rinha.json"]
