FROM golang:1.22.1

WORKDIR /app

copy go.mod .
copy main.go .
copy helloworld.go .

RUN go build -o bin .

ENTRYPOINT ["/app/bin"]

