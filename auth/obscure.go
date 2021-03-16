package auth

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
