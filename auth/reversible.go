package auth

import "github.com/imakiri/playground/core"

type JWT struct {
	Reversible
}

func (t JWT) New(id core.ID) (core.Credential, error) {
	var cred core.Credential
	var err error

	// TODO: JWT.NewFormID

	return cred, err
}

func (t JWT) GetID(credential core.Credential) (core.ID, error) {
	var id core.ID
	var err error

	// TODO: JWT.GetID

	return id, err
}
