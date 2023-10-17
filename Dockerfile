FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/

FROM alpine:latest
WORKDIR /app


COPY --from=build /app .



# Copy the built Go binary from the build stage to the runtime stage

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]