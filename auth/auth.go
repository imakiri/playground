package auth

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
)

//type Storage interface {
//	Create(pemid utils.PEMID, credential utils.Credential) (utils.UUID, error)
//	Read(credential utils.Credential) (utils.UUID, utils.PEMID, error)
//	ReadCred(uuid utils.UUID) (utils.PEMID, utils.Credential, error)
//	Update(uuid utils.UUID, pemid *utils.PEMID, credential *utils.Credential) error
//	Delete(uuid utils.UUID) error
//}

//type Randomizer interface {
//	Random() service.CredentialObscure
//}
//
//type Hasher interface {
//	Hash(plain service.CredentialPlain) service.CredentialObscure
//	Compare(obscure0, obscure1 service.CredentialObscure) bool
//}
//
//type StrongAuthenticator interface {
//	Validate(plain service.CredentialPlain) (service.CredentialObscure, error)
//	Authenticate(valid service.CredentialObscure, key service.CredentialPlain) (service.ID, error)
//}
//
//type WeakAuthenticator interface {
//	Authenticate(obscure service.CredentialObscure) (service.ID, error)
//}
//
//type StrongRegistrar interface {
//	Register(plain service.CredentialPlain) (service.CredentialObscure, error)
//	Activate(valid service.CredentialObscure, key service.CredentialPlain) (service.ID, error)
//	Revoke(id service.ID) error
//}
//
//type WeakRegistrar interface {
//	Register(id service.ID, expireAt time.Time) (service.CredentialObscure, error)
//	Revoke(id service.ID) error
//}
//
//type Strong interface {
//	StrongRegistrar
//	StrongAuthenticator
//}
//
//type Weak interface {
//	WeakRegistrar
//	WeakAuthenticator
//}

type Service struct {
	service.Service
	config *cfg.Auth
}

func New(bs service.Service) (*Service, error) {
	var s Service
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4Auth(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}

//func (s Service) Authenticate(credentials []service.Credential) (service.ID, error) {
//	var id service.ID
//	var err error
//
//	// TODO: Implement Authenticate
//
//	return id, err
//}
