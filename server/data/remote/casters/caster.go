package casters

import (
	"github.com/imakiri/playground/server/core"
	"sync"
)

type Caster interface {
	Cast(group *sync.WaitGroup, c chan core.ThingImp)
}
