package interfaces

import "net/http"

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
	Channel        chan Thing
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

type App interface {
	Do0(parcel Parcel)
	Do1(parcel Parcel)
	Do2(parcel Parcel)
	Do3(parcel Parcel)
	Do4(parcel Parcel)
}

type Resolver func(p Parcel)
