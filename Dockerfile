FROM golang:1.21 as build

WORKDIR /go/src/app

RUN apt-get update -y && \
    apt-get install -y libwebp-dev ffmpeg

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/app /

COPY --from=build /usr/bin/ffmpeg /usr/bin/ffmpeg
COPY --from=build /usr/lib/x86_64-linux-gnu /usr/lib/x86_64-linux-gnu

EXPOSE 8080

CMD ["/app"]