FROM amd64/golang:1.15-alpine AS build-env

ENV GO111MODULE=auto
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk add bash ca-certificates git gcc libc-dev

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal/asset/transport internal/asset/transport
COPY internal/web internal/web
COPY internal/utils internal/utils
COPY cmd cmd

RUN go build -o launcher ./cmd/web/main.go

FROM amd64/ubuntu:bionic AS run-env

RUN apt update && apt install libc-dev -y && apt install musl -y

WORKDIR /srv

COPY --from=build-env /src/launcher .

CMD ./launcher -debug=false