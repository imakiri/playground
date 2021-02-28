package auth

import (
	"github.com/imakiri/playground/admin/cfg"
	"github.com/imakiri/playground/core"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   *cfg.Auth
	configDB *cfg.Data
}

func NewService(c *cfg.System) (*Service, error) {
	var s Service
	var err error

	//var _ internal.Cookie

	s.config = c.GetAuth()
	s.configDB = c.GetData()

	s.db, err = core.Connect(c.GetData())
	if err != nil {
		return nil, err
	}

	return &s, err
}

type Action interface{}
type Factor interface{}
type ID interface{}

type Resolver interface {
	Resolve([]Factor) (ID, error)
}

type Registrar interface {
	Enrol(Factor) (bool, error)
}

type Authenticator interface {
	Check(Factor) (bool, error)
	ID() (ID, error)
}

type Authorizer interface {
	Permit(ID, Action) (bool, error)
}
