package auth

import (
	"github.com/imakiri/gorum/core"
	"github.com/imakiri/gorum/data"
	"github.com/imakiri/gorum/transport"
	"github.com/jackc/pgx/v4"
	"time"
)

type Randomizer interface {
	Random() core.CredentialObscure
}

type Hasher interface {
	Hash(plain core.CredentialPlain) core.CredentialObscure
	Compare(obscure0, obscure1 core.CredentialObscure) bool
}

type StrongAuthenticator interface {
	Validate(plain core.CredentialPlain) (core.CredentialObscure, error)
	Authenticate(valid core.CredentialObscure, key core.CredentialPlain) (core.ID, error)
}

type WeakAuthenticator interface {
	Authenticate(obscure core.CredentialObscure) (core.ID, error)
}

type StrongRegistrar interface {
	Register(plain core.CredentialPlain) (core.CredentialObscure, error)
	Activate(valid core.CredentialObscure, key core.CredentialPlain) (core.ID, error)
	Revoke(id core.ID) error
}

type WeakRegistrar interface {
	Register(id core.ID, expireAt time.Time) (core.CredentialObscure, error)
	Revoke(id core.ID) error
}

type Strong interface {
	StrongRegistrar
	StrongAuthenticator
}

type Weak interface {
	WeakRegistrar
	WeakAuthenticator
}

func NewService(c *transport.System) (*Service, error) {
	var s Service
	var err error

	s.config = c.GetAuth()
	s.configDB = c.GetData()
	s.db, err = data.Connect(c.GetData())
	if err != nil {
		return nil, err
	}

	return &s, err
}

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   *transport.Auth
	configDB *transport.Data
}

func (s Service) Authenticate(credentials []core.Credential) (core.ID, error) {
	var id core.ID
	var err error

	// TODO: Implement Authenticate

	return id, err
}
