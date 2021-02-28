package core

import (
	"github.com/imakiri/playground/data"
)

// Auth

type AuthKey []byte

// Request

type AuthRequestLogin struct {
	Login    data.ModelCredentialsLogin
	Password string
}
type AuthRequestCheckAccess struct {
	Key AuthKey
}
type AuthRequestLogout struct {
	Key AuthKey
}

//

// Response

type AuthResponseLogin struct {
	Meta
	Key AuthKey
}
type AuthResponseLogout struct {
	Meta
}

//

type Action interface{}
type Factor interface{}
type Factors []Factor
type ID interface{}

type Resolver interface {
	Resolve(Factors) (ID, error)
}

type Registrar interface {
	Enrol(Factor) (bool, error)
}

type Authenticator interface {
	Check(Factor) (bool, error)
	ID() (ID, error)
}

type Authorizer interface {
	Permit(ID, Action) (bool, error)
}

type Storage interface {
	Read(Factors) (Factors, error)
	Write(Factors) error
}

//
