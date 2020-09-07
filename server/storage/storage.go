package storage

import (
	"github.com/imakiri/playground/server/core"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Thing struct {
	Header string
	Data   []byte
	Error  error
}

func (th *Thing) GetHeader() string {
	return th.Header
}

func (th *Thing) GetData() []byte {
	return th.Data
}

func (th *Thing) GetError() error {
	return th.Error
}

type Local bool

func (l *Local) GetThing(str string, c chan core.Thing) {

}

func (l *Local) DoThing(str string, c chan core.Thing) {

}

func (l *Local) ChangeThing(str string, th core.Thing, c chan core.Thing) {

}

func (l *Local) StoreThing(str string, th core.Thing, c chan core.Thing) {

}
