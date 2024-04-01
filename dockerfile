FROM golang:1.22.1

WORKDIR /usr/src/app

copy . .
RUN go mod tidy