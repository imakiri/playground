package auth

import (
	"github.com/imakiri/playground/core"
)

func NewCookie(encoder Encoder, storage Irreversible) (*Cookie, error) {
	var c Cookie
	var err error

	c.encoder = encoder
	c.storage = storage

	return &c, err
}

type Cookie struct {
	encoder Encoder
	storage Irreversible
}

func (c Cookie) New(id core.ID) (core.Credential, error) {
	var cred core.Credential
	var err error

	// TODO: Cookie.NewFormID

	return cred, err
}

func (c Cookie) GetID(credential core.Credential) (core.ID, error) {
	var id core.ID
	var err error

	// TODO: Cookie.GetID

	return id, err
}
