package internal

import (
	"github.com/imakiri/gorum/core"
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

func (e random) Random() core.CredentialObscure {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, e.length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}
