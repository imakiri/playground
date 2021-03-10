package sender

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/erres"
	"reflect"
)

type address interface {
	Address() []string
}

type sender interface {
	Send(addr address, msg []byte) error
}

type constructor interface {
	NewAddress(addr ...string) (address, error)
}

type Service interface {
	core.Service
	constructor
	sender
}

const N_Sender core.ServiceName = "Service"

func NewService() (*service, error) {
	var s service
	var err error

	s.email, err = newServiceEmailBasic("", "", "", "", nil)

	// TODO: Sender service constructor

	return &s, err
}

type service struct {
	email serviceEmail
}

func (service) Name() core.ServiceName {
	return N_Sender
}

func (s service) Send(addr address, msg []byte) error {
	switch a := addr.(type) {
	case email:
		return s.email.Send(msg, a)
	default:
		return erres.
			E_TypeMismatch.
			Time("").
			Route(s.Name().String()).
			Description(reflect.TypeOf(a).String())
	}
}
