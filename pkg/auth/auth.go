package auth

import "github.com/imakiri/playground/core"

type Auth struct {
	dataService core.DataService
	salt        []byte
}

func NewAuthService(s core.Settings) (*Auth, error) {
	var auth *Auth

	return auth, nil
}
