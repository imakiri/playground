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

type Judge interface {
	// Judge register new Assertion taking into account Assertions made before and return Assertion back along with error
	AddAssertion(Assertion, Assertions) (Assertion, error)

	// Judge check Assertion taking into account Assertions made before and return Assertion back along with error
	CheckAssertion(Assertion, Assertions) (Assertion, error)
}

type Credentials interface {
	Judges() []Judge
	ID() string
	Level() int8
	IsVerified() bool

	RegisterAssertion(a Assertion, j Judge) error
	VerifyAssertion(a Assertion, j Judge) error
	WithdrawAssertion() error
}

type Assertion interface {
	Type() string
	Data() interface{}
}
type Assertions []Assertion

//
