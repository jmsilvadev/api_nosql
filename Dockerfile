FROM golang:latest as build

LABEL version="0.1.0"

WORKDIR /go/src/api-nosql