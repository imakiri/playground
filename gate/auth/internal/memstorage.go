package internal

import (
	"github.com/imakiri/playground/core"
	"google.golang.org/grpc/codes"
	"math/rand"
	"sync"
)

type MemStorage struct {
	data  map[Assertion_Rand]Assertion_ID
	rwmux *sync.RWMutex
}

func (mem *MemStorage) AddAssertion(id core.Assertion, _ core.Assertions) (core.Assertion, error) {
	switch id := id.(type) {
	case Assertion_ID:
		var r Assertion_Rand
		r = Assertion_Rand(random(60))

		mem.rwmux.Lock()
		defer mem.rwmux.Unlock()
		mem.data[r] = id

		return r, nil
	default:
		return nil, core.StatusCode(codes.InvalidArgument)
	}
}

func (mem *MemStorage) CheckAssertion(assertion core.Assertion, assertions core.Assertions) (core.Assertion, error) {
	panic("implement me")
}

func random(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
