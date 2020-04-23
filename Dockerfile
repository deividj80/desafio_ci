FROM golang:1.6.1-alpine

WORKDIR  /go/src/go-sum

COPY ./src/main/ /go/src/go-sum