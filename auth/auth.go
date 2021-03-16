package auth

import (
	"github.com/imakiri/gorum/data"
	"github.com/imakiri/gorum/utils"
	"github.com/jackc/pgx/v4"
	"time"
)

type Randomizer interface {
	Random() utils.CredentialObscure
}

type Hasher interface {
	Hash(plain utils.CredentialPlain) utils.CredentialObscure
	Compare(obscure0, obscure1 utils.CredentialObscure) bool
}

type StrongAuthenticator interface {
	Validate(plain utils.CredentialPlain) (utils.CredentialObscure, error)
	Authenticate(valid utils.CredentialObscure, key utils.CredentialPlain) (utils.ID, error)
}

type WeakAuthenticator interface {
	Authenticate(obscure utils.CredentialObscure) (utils.ID, error)
}

type StrongRegistrar interface {
	Register(plain utils.CredentialPlain) (utils.CredentialObscure, error)
	Activate(valid utils.CredentialObscure, key utils.CredentialPlain) (utils.ID, error)
	Revoke(id utils.ID) error
}

type WeakRegistrar interface {
	Register(id utils.ID, expireAt time.Time) (utils.CredentialObscure, error)
	Revoke(id utils.ID) error
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
	log      utils.LogService
	config   *transport.Auth
	configDB *transport.Data
}

func (s Service) Authenticate(credentials []utils.Credential) (utils.ID, error) {
	var id utils.ID
	var err error

	// TODO: Implement Authenticate

	return id, err
}
