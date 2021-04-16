package data

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/types"
	"github.com/imakiri/gorum/utils"
	"github.com/jmoiron/sqlx"
)

const path = "secrets/data.yaml"

type service struct {
	secret types.SecretData
	db     *sqlx.DB
}

func newService() (*service, error) {
	var secret types.SecretData
	var err = utils.ReadYAML(path, &secret)
	if err != nil {
		return nil, erres.InternalServiceError.Extend(0)
	}

	var s service
	s.secret = secret
	if s.db, err = sqlx.Connect("pgx", secret.Postgres.DSN); err != nil {
		return nil, erres.ConnectionError.Extend(0)
	}
	return &s, err
}

func NewUserService() (User, error) {
	var err error
	var user user
	if user.service, err = newService(); err != nil {
		return nil, err
	}

	return user, err
}

func NewCookieService() (Cookie, error) {
	var err error
	var cookie cookie
	if cookie.service, err = newService(); err != nil {
		return nil, err
	}

	return cookie, err
}

func NewLogpassService() (Logpass, error) {
	var err error
	var logpass logpass
	if logpass.service, err = newService(); err != nil {
		return nil, err
	}

	return logpass, err
}

func NewThreadService() (Thread, error) {
	var err error
	var thread thread
	if thread.service, err = newService(); err != nil {
		return nil, err
	}

	return thread, err
}

func NewPostService() (Post, error) {
	var err error
	var post post
	if post.service, err = newService(); err != nil {
		return nil, err
	}

	return post, err
}
