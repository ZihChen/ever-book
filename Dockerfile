FROM golang:1.17-alpine

RUN apk add git \
    && go get github.com/pilu/fresh \