# syntax=docker/dockerfile:1

FROM golang:latest
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN mkdir -p ./config
COPY config/* ./config
RUN mkdir -p ./details
COPY details/* ./details
RUN mkdir -p ./docs
COPY docs/* ./docs
COPY local-docker-env/env.json ./
COPY main.go ./

RUN go mod download

RUN go build -o /product-details

EXPOSE 8080

CMD ["/product-details"]
