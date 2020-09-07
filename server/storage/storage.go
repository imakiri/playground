package storage

import (
	"github.com/imakiri/playground/server/core"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Local bool

func (l *Local) GetThing(str string, c chan core.Thing) {

}

func (l *Local) DoThing(str string, c chan core.Thing) {

}

func (l *Local) ChangeThing(str string, th core.Thing, c chan core.Thing) {

}

func (l *Local) StoreThing(str string, th core.Thing, c chan core.Thing) {

}
