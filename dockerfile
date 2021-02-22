FROM golang:alpine  AS build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc libc-dev

WORKDIR /playground
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o server ./cmd/mainServiceLauncher/main.go

CMD ./server
