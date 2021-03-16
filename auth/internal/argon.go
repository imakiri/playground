package internal

import (
	"bytes"
	"github.com/imakiri/gorum/utils"
	"golang.org/x/crypto/argon2"
)

const Argon2i argonType = "argon2i"
const Argon2id argonType = "argon2id"

type argonType string

func NewArgon(t argonType, iterations uint32, threads uint8, keyLen uint32, salt []byte, memory uint32) (*argon, error) {
	var arg2 argon
	arg2._type = t
	arg2.iterations = iterations
	arg2.threads = threads
	arg2.keyLen = keyLen
	arg2.salt = salt
	arg2.memory = memory

	return &arg2, nil
}

type argon struct {
	_type      argonType
	iterations uint32
	threads    uint8
	keyLen     uint32
	salt       []byte
	memory     uint32
}

func (e argon) Hash(plain utils.CredentialPlain) utils.CredentialObscure {
	var obscure utils.CredentialObscure

	switch e._type {
	case Argon2i:
		obscure = argon2.Key([]byte(plain), e.salt, e.iterations, e.memory, e.threads, e.keyLen)
	case Argon2id:
		obscure = argon2.IDKey([]byte(plain), e.salt, e.iterations, e.memory, e.threads, e.keyLen)
	default:
		return nil
	}

	return obscure
}

func (e argon) Compare(obscure0, obscure1 utils.CredentialObscure) bool {
	return bytes.Equal(obscure0, obscure1)
}
