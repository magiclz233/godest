FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/godest ./cmd/server

FROM alpine:3.21

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app
COPY --from=builder /out/godest /app/godest
COPY config/config.yaml /app/config/config.yaml

USER app
EXPOSE 8080

ENTRYPOINT ["/app/godest"]
