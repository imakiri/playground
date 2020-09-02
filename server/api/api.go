package api

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
