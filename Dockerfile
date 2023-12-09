# Start by building the application.
FROM golang:1.21 as build

WORKDIR /go/src/app

# webpに変換する用
RUN  apt-get update -y
RUN apt-get install libwebp-dev -y

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /

EXPOSE 8080

USER yuorei:yuorei-group

CMD ["/app"]