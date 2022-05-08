FROM golang:1.15-alpine as dev
WORKDIR /workdir

FROM golang:1.15-alpine as build
WORKDIR /videos
COPY . /videos
RUN go build -o videos

FROM alpine as runtime
COPY --from=build /videos/videos /usr/local/bin/videos
COPY ./videos.json /
COPY run.sh /
ENTRYPOINT ["./run.sh"]
