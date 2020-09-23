package core

import (
	"net/http"
)

type Thing interface {
	GetHeader() string
	GetData() []byte
	GetError() error
}

type Parcel struct {
	Channel        *chan ThingImp
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
