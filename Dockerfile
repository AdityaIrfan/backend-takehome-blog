FROM golang:1.21.0

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest

COPY ./app .

EXPOSE 8080

CMD air