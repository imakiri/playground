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
	AddAssertion(rawAss Assertion, c Credentials) (Assertion, error)
	CheckAssertion(rawAss Assertion, c Credentials) (Assertion, error)
	WithdrawAssertion(ass Assertion, c Credentials) error
}
type Judges []Judge

type Credentials interface {
	Judges() Judges
	Assertions() Assertions
	ID() Assertion
	Level() int
	IsVerified() bool

	ExtendWith(c Credentials) error
	RegisterAssertion(a Assertion) error
	VerifyAssertion(a Assertion) error
	WithdrawAssertion() error
}

type Assertion interface{}
type Assertions []Assertion

type Storage interface {
	Read(...Assertion) (Assertions, error)
	Write(...Assertion) error
}

//
