package core

import (
	"net/http"
)

type Thing interface {
	GetHeader() string
	GetData() []byte
	GetError() error
}

type Api interface {
	GetThing(str string, c chan Thing)
	DoThing(str string, c chan Thing)
	ChangeThing(str string, th Thing, c chan Thing)
	StoreThing(str string, th Thing, c chan Thing)
}

type Parcel struct {
	Channel        *chan ThingImp
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
