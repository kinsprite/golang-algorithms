# build
FROM golang:1.13-alpine as build

ENV PORT 8080
EXPOSE 8080

RUN mkdir /app
ADD . /app

ENV GOPROXY https://goproxy.io
ENV GIN_MODE release

WORKDIR  /app
RUN go mod vendor
RUN go build -mod=vendor -o golang-algorithms .


# release
FROM alpine:3.10
RUN mkdir /app
COPY --from=build /app/golang-algorithms /app/golang-algorithms

WORKDIR  /app
CMD ["/app/golang-algorithms"]
