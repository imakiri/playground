package auth

import (
	"context"
	"github.com/imakiri/gorum/internal/cfg"
	"google.golang.org/grpc"
)

//type Storage interface {
//	Create(pemid utils.PEMID, credential utils.Credential) (utils.PostUUID, error)
//	Read(credential utils.Credential) (utils.PostUUID, utils.PEMID, error)
//	ReadCred(uuid utils.PostUUID) (utils.PEMID, utils.Credential, error)
//	Update(uuid utils.PostUUID, pemid *utils.PEMID, credential *utils.Credential) error
//	Delete(uuid utils.PostUUID) error
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

type Config interface {
	Get4Auth(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.Auth, error)
}

type Service struct {
	config       Config
	configCached *cfg.Auth
}

func NewService(c Config) (*Service, error) {
	var s Service
	var err error

	s.config = c
	s.configCached, err = s.config.Get4Auth(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
