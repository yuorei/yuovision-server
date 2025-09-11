FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o yuovision-server

FROM debian:bookworm-slim

# Install curl for health checks
RUN apt-get update && apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Create config directory
RUN mkdir -p /app/config

COPY --from=build /app/yuovision-server .

EXPOSE 8080

ENV TZ=Asia/Tokyo
ENV HTTP_PORT=8080

# Create non-root user
RUN groupadd -r appuser && useradd -r -g appuser appuser
RUN chown -R appuser:appuser /app
USER appuser

CMD ["./yuovision-server"]