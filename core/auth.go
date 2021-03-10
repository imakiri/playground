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
type ID interface {
	UUID() uint64
	PemID() uint64
}

type UUID uint64

type PEMID uint64

type Key interface{}

type Factor interface{}

type Credentials struct {
	Key    Key
	Factor Factor
}

type Validator interface {
	Validate(Credentials) error
}

type Storage interface {
	Create(pemid PEMID, credentials Credentials) (UUID, error)
	Read(credentials Credentials) (UUID, PEMID, error)
	ReadCred(uuid UUID) (PEMID, Credentials, error)
	Update(uuid UUID, pemid *PEMID, credentials *Credentials) error
	Delete(uuid UUID) error
}

type Registrar interface {
	Register(ID) Key
}

type Addressee interface {
	NewAddress(string) (string, error)
}

type Hasher interface {
	Hash(raw []byte) []byte
	Compare(raw, hashed []byte) bool
}

type Sender interface {
	Send(addr Addressee, msg []byte) error
}

type Resolver interface {
	Resolve(Factor) (ID, error)
}

type Authenticator interface {
	Authenticate(ID, Factor) (ID, error)
}

type V interface {
	V(ID, Action) (bool, error)
}

//type Authority interface {
//	NewID() (ID, error)
//	ResetID(ID) error
//}

//type Gateway interface {
//	Controllers() (Validator, Sender, Decoder, Authenticator, V)
//}
//
//type Gate interface {
//	Gateways() []Gateway
//}

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
