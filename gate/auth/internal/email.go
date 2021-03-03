package internal

import (
	"bytes"
	"errors"
	"github.com/imakiri/playground/core"
	"net/smtp"
	"text/template"
)

func NewEmailService(addr, user, password, from string, rand func() string, template *template.Template) (*Email, error) {
	var email Email
	var err error

	email.auth = smtp.PlainAuth("", user, password, addr)
	email.from = from
	email.rand = rand
	email.temp = template

	return &email, err
}

type Email struct {
	temp *template.Template
	addr string
	auth smtp.Auth
	from string
	rand func() string
}

func (e Email) Encode(key core.Key) (core.Factor, error) {
	var err error
	switch k := key.(type) {
	case Key_Email:
		var code = Factor_Code(e.rand())
		var to []string
		to = append(to, string(k))

		var msg_buf = bytes.NewBuffer([]byte{})

		err = e.temp.Execute(msg_buf, code)
		if err != nil {
			return nil, errors.New("template error")
		}

		err = smtp.SendMail(e.addr, e.auth, e.from, to, msg_buf.Bytes())
		if err != nil {
			return nil, err
		}

		return code, err
	default:
		return err, errors.New("wrong key type")
	}
}
