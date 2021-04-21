package auth

const (
	Login typePlain = "login"
	Email typePlain = "email"
	SMS   typePlain = "sms"
)

type typePlain string

type Plain interface {
	Type() typePlain
	Identifier() string
	Key() string
}

// Login ---------------------------------------------------------------------------------------------------------------

type login struct {
	_type    typePlain
	name     string
	password string
}

func (l login) Type() typePlain {
	return l._type
}

func (l login) Identifier() string {
	return l.name
}

func (l login) Key() string {
	return l.password
}

// Email ---------------------------------------------------------------------------------------------------------------

type email struct {
	_type   typePlain
	address string
	code    string
}

func (e email) Type() typePlain {
	return e._type
}

func (e email) Identifier() string {
	return e.address
}

func (e email) Key() string {
	return e.code
}

// SMS -----------------------------------------------------------------------------------------------------------------

type sms struct {
	_type  typePlain
	number string
	code   string
}

func (s sms) Type() typePlain {
	return s._type
}

func (s sms) Identifier() string {
	return s.number
}

func (s sms) Key() string {
	return s.code
}
