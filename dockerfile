FROM golang:1.22.1

WORKDIR /app

copy go.mod .
copy main.go .
copy index.html .


RUN go build -o bin .

ENTRYPOINT ["/app/bin"]

