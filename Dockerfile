FROM golang:1.21 as build

WORKDIR /go/src/app

RUN apt-get update -y && \
    apt-get install -y libwebp-dev ffmpeg

COPY . .

RUN go mod download

# 本番環境用で運用する際には修正する
ENV AWS_ACCESS_KEY_ID=minioadmin
ENV AWS_SECRET_ACCESS_KEY=minioadmin
ENV AWS_S3_ENDPOINT=http://minio:9000
ENV MONGODB_URI=mongodb://root:example@mongo:27017/

RUN  go build -o /app

EXPOSE 8080

CMD ["/app"]