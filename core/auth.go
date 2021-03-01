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
type Key interface{}
type Factor interface{}

type Encoder interface {
	Encode(Key) (Factor, error)
}

type Decoder interface {
	Decode(Factor) (Key, error)
}

type Resolver interface {
	Resolve(Key) (ID, error)
}

type Validator interface {
	Validate(ID, Action) (bool, error)
}

//type Authority interface {
//	NewID() (ID, error)
//	ResetID(ID) error
//}

type Gate interface {
}

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
