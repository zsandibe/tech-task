FROM golang:alpine AS build
ENV CGO_ENABLED=1
RUN apk add --no-cache \
    gcc \
    musl-dev
WORKDIR /app
COPY . .
RUN go build -o task ./cmd/main.go
FROM alpine:latest
WORKDIR /app
COPY --from=build /app .
CMD ["./task"]