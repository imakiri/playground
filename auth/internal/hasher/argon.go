package hasher

import (
	"errors"
	"github.com/imakiri/playground/auth/internal"
	"github.com/imakiri/playground/core"
	"golang.org/x/crypto/argon2"
)

const Argon2i = "argon2i"
const Argon2id = "argon2id"

type Argon2Type string

func NewArgon2(t Argon2Type, iterations uint32, threads uint8, keyLen uint32, salt []byte, memory uint32) (*Argon2, error) {
	var arg2 Argon2
	arg2._type = t
	arg2.iterations = iterations
	arg2.threads = threads
	arg2.keyLen = keyLen
	arg2.salt = salt
	arg2.memory = memory

	return &arg2, nil
}

type Argon2 struct {
	_type      Argon2Type
	iterations uint32
	threads    uint8
	keyLen     uint32
	salt       []byte
	memory     uint32
}

func (e Argon2) Send(key core.Key) (core.Factor, error) {
	switch k := key.(type) {
	case internal.Key_Password:
		var hash internal.Factor_Hash

		switch e._type {
		case Argon2i:
			hash = argon2.Key([]byte(k), e.salt, e.iterations, e.memory, e.threads, e.keyLen)
		case Argon2id:
			hash = argon2.IDKey([]byte(k), e.salt, e.iterations, e.memory, e.threads, e.keyLen)
		}

		return hash, nil
	default:
		return nil, errors.New("wrong key type")
	}
}
