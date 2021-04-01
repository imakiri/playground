FROM amd64/golang:1.15-alpine AS build-env

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk add bash ca-certificates git gcc libc-dev

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./types ./types
COPY ./transport ./transport
COPY ./assets ./assets
COPY ./web ./web
COPY ./cmd ./cmd
COPY ./secrets ./secrets

RUN go build -o run ./cmd/web/main.go

FROM amd64/ubuntu:bionic AS run-env

RUN apt update && apt install libc-dev -y && apt install musl -y

WORKDIR /server

COPY --from=build-env /src .

CMD ./run