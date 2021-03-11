package internal

import (
	"github.com/imakiri/playground/core"
	"math/rand"
)

func NewRandom(length uint) (*random, error) {
	var r random
	var err error

	r.length = length
	return &r, err
}

type random struct {
	length uint
}

func (e random) Encode(core.ID) core.Credential {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, e.length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return core.Credential(b)
}
