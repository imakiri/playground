package internal

import (
	"github.com/imakiri/playground/core"
	"sync"
)

type MemStorage struct {
	data map[Assertion_Rand]struct {
		ID         Assertion_ID
		ExpireTime Assertion_ExpirationTime
	}
	rwmux *sync.RWMutex
}

func (mem *MemStorage) Read(assertion ...core.Assertion) (core.Assertions, error) {
	mem.rwmux.RLock()
	defer mem.rwmux.RUnlock()

	panic("implement me")
}

func (mem *MemStorage) Write(assertion ...core.Assertion) error {
	mem.rwmux.Lock()
	defer mem.rwmux.Unlock()

	panic("implement me")
}
