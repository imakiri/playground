package message

import (
	"github.com/imakiri/erres"
)

type _type string

const Email _type = "email"

type Messenger interface {
	Send(msg []byte) error
}

type messengerEmail struct {
	service *ServiceEmail
	addr    email
}

func (e messengerEmail) Send(msg []byte) error {
	return e.service.Send(e.addr, msg)
}

type Service interface {
	New(p _type, addr []string) (Messenger, error)
}

func NewService(features ..._type) (*service, error) {
	var s service
	var err error

	s.email, err = NewServiceEmail("", "", "", "", nil)
	if err != nil {
		return nil, err
	}

	for _, v := range features {
		s.features[v] = true
	}

	return &s, err
}

type service struct {
	features map[_type]bool
	email    *ServiceEmail
}

func (s service) New(p _type, addr []string) (Messenger, error) {
	switch p {
	case Email:
		if _, ok := s.features[p]; !ok {
			return nil, erres.NotFound.SetTime("")
		}

		var m messengerEmail
		var err error

		m.service = s.email
		m.addr, err = NewEmail(addr...)
		if err != nil {
			return nil, err
		} else {
			return m, err
		}
	default:
		return nil, erres.TypeMismatch.SetTime("")
	}
}
