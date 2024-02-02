FROM golang:1.21 as build

WORKDIR /go/src/app

RUN apt-get update -y && \
    apt-get install -y libwebp-dev ffmpeg

COPY . .

RUN go mod download

RUN  go build -o /app

EXPOSE 8080

CMD ["/app"]