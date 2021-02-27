package storage

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/gate/auth/internal"
	"sync"
)

type MemStorage struct {
	data map[internal.Assertion_Rand]struct {
		ID         internal.Assertion_ID
		ExpireTime internal.Assertion_ExpirationTime
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
