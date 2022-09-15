# syntax=docker/dockerfile:1

## Build
FROM golang:alpine AS builder
WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN mkdir -p ./config
COPY config/* ./config
RUN mkdir -p ./details
COPY internal/details/* ./details
RUN mkdir -p ./docs
COPY api/docs/* ./docs

COPY main.go ./

RUN go mod download

RUN go build -o /product-details

## Deploy
FROM alpine:latest

WORKDIR /


COPY --from=builder /bin/app/product-details /product-details
COPY configs/local-docker-env/env.json ./

EXPOSE 8080



ENTRYPOINT ["/product-details"]

