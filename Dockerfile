# syntax=docker/dockerfile:1

FROM golang:latest
WORKDIR /app
RUN go mod init github.com/MihasBel/product-details
COPY go.mod ./
COPY go.sum ./
COPY config/* ./config
COPY details/* ./details
COPY docs/* ./docs
COPY main.go ./

RUN go mod download

RUN go build github.com/MihasBel/product-details -o /product-details

EXPOSE 8080

CMD ["/product-details"]
