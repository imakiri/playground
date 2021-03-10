package sender

import (
	"bytes"
	"github.com/imakiri/playground/core"
	error2 "github.com/imakiri/playground/erres"
	"net/smtp"
	"text/template"
)

type serviceEmail interface {
	core.Service
	Send(msg []byte, addr email) error
	NewAddress(addr ...string) (*email, error)
}

type email []string

func (e *email) setAddress(addr ...string) {
	*e = addr
}

func (e email) Address() []string {
	return e
}

const N_BasicEmailService core.ServiceName = "Basic Email Service"

func newServiceEmailBasic(addr, user, password, from string, template *template.Template) (*serviceEmailBasic, error) {
	var err error
	if err = core.Validate(addr).AsEmailAddress(); err != nil {
		return nil, err
	}

	var email serviceEmailBasic

	email.auth = smtp.PlainAuth("", user, password, addr)
	email.from = from
	email.temp = template

	return &email, err
}

type serviceEmailBasic struct {
	temp *template.Template
	addr string
	auth smtp.Auth
	from string
}

func (serviceEmailBasic) Name() core.ServiceName {
	return N_BasicEmailService
}

func (e serviceEmailBasic) Send(msg []byte, addr email) error {
	var err error
	var msg_buf = bytes.NewBuffer([]byte{})

	err = e.temp.Execute(msg_buf, msg)
	if err != nil {
		return error2.E_DeserializationError.Description(err.Error())
	}

	err = smtp.SendMail(e.addr, e.auth, e.from, addr.Address(), msg_buf.Bytes())
	if err != nil {
		return error2.E_InternalServiceError.Description(err.Error())
	}

	return err
}

func (serviceEmailBasic) NewAddress(addr ...string) (*email, error) {
	var e = new(email)
	var checked_addr []string
	var err error

	for _, a := range addr {

		// TODO: Need to check if addr is correct email address

		checked_addr = append(checked_addr, a)
	}

	e.setAddress(checked_addr...)
	return e, err
}
