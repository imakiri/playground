package internal

import (
	"github.com/imakiri/erres"
)

const (
	Temporary typeObscure = "temporary"
	Cookie    typeObscure = "cookie"
	JWT       typeObscure = "jwt"
)

type typeObscure string

type Obscure interface {
	Type() typeObscure
	Key() []byte
}

func NewObscure(t typeObscure, key []byte) (Obscure, error) {
	var o Obscure
	var err error

	switch t {
	case Temporary:
		o = temporary{
			_type: Temporary,
			data:  key,
		}
	case Cookie:
		o = cookie{
			_type: Cookie,
			data:  key,
		}
	case JWT:
		o = jwt{
			_type: JWT,
			data:  key,
		}
	default:
		err = erres.TypeMismatch.SetTime("")
	}

	return o, err
}

// Temporary -----------------------------------------------------------------------------------------------------------

type temporary struct {
	_type typeObscure
	data  []byte
}

func (t temporary) Type() typeObscure {
	return t._type
}

func (t temporary) Key() []byte {
	return t.data
}

// Cookie --------------------------------------------------------------------------------------------------------------

type cookie struct {
	_type typeObscure
	data  []byte
}

func (c cookie) Type() typeObscure {
	return c._type
}

func (c cookie) Key() []byte {
	return c.data
}

// JWT -----------------------------------------------------------------------------------------------------------------

type jwt struct {
	_type typeObscure
	data  []byte
}

func (j jwt) Type() typeObscure {
	return j._type
}

func (j jwt) Key() []byte {
	return j.data
}
