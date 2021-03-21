package message

import (
	"bytes"
	"github.com/imakiri/erres"
	"net/smtp"
	"text/template"
)

func isEmail(addr string) bool {

	// TODO: Email validator

	return true
}

func NewEmail(addr ...string) (email, error) {
	var e email
	var err error

	for _, a := range addr {
		if isEmail(a) {
			e = append(e, a)
		} else {
			return nil, erres.TypeMismatch.SetRoute("message").SetRoute("NewEmail").SetTime("")
		}
	}

	return e, err
}

type email []string

func (e email) Address() []string {
	return e
}

func NewServiceEmail(addr, user, password, from string, template *template.Template) (*ServiceEmail, error) {
	var err error
	if !isEmail(addr) {
		return nil, erres.TypeMismatch.SetRoute("message").SetRoute("NewServiceEmail").SetDescription("addr").SetDescription(addr).SetTime("")
	}

	var email ServiceEmail

	email.auth = smtp.PlainAuth("", user, password, addr)
	email.from = from
	email.template = template

	return &email, err
}

type ServiceEmail struct {
	template *template.Template
	addr     string
	auth     smtp.Auth
	from     string
}

func (e ServiceEmail) Send(addr email, msg []byte) error {
	var err error
	var msg_buf = bytes.NewBuffer([]byte{})

	err = e.template.Execute(msg_buf, msg)
	if err != nil {
		return erres.DeserializationError.SetDescription(err.Error())
	}

	err = smtp.SendMail(e.addr, e.auth, e.from, addr.Address(), msg_buf.Bytes())
	if err != nil {
		return erres.InternalServiceError.SetDescription(err.Error())
	}

	return err
}
