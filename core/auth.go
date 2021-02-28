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
type ID interface{}

type Inspector interface {
	Enrol(...Factor) (bool, error)
	Withdraw(...Factor) (bool, error)
}

type Verifier interface {
	Verify(...Factor) (ID, error)
}

type Resolver interface {
	Resolve(factor Factor)
}

type Worker interface {
	GetResolvers() []Resolver
	GetID() (ID, error)
}

type Identificator interface {
	Worker
	Identify(Factor) (bool, error)
	Withdraw(Factor) (bool, error)
}

type Authenticator interface {
	Worker
	Check(Factor) (bool, error)
}

type Authorizer interface {
	Worker
	Permit(ID, Action) (bool, error)
}

//

func IsNilSafe(l ...interface{}) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			return false
		}
	}
	return true
}

func IsNilSafeEx(l ...interface{}) (b []bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			b = append(b, false)
		} else {
			b = append(b, true)
		}
	}
	return
}
