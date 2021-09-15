FROM golang:1.16-alpine as builder

COPY . /app/throttle
WORKDIR /app/throttle

RUN apk add --no-cache nano bash shadow build-base

RUN mkdir /.cache
RUN chown nobody:nobody -R /.cache

USER nobody