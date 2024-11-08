FROM golang:1.23 as build

WORKDIR /go/src/app

RUN apt-get update -y && \
    apt-get install -y libwebp-dev ffmpeg

COPY . .

RUN go mod download

RUN  go build -o /app
ENV TZ Asia/Tokyo

EXPOSE 50051

CMD ["/app"]