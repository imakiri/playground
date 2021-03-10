package auth

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/core"
	"github.com/jackc/pgx/v4"
)

type Credential []byte

type ID struct {
	uuid  uint64
	pemid []uint64
}

func (e ID) UUID() uint64 {
	return e.uuid
}

func (e ID) PemID() []uint64 {
	return e.pemid
}

func NewService(c *cfg.System) (*Service, error) {
	var s Service
	var err error

	s.config = c.GetAuth()
	s.configDB = c.GetData()

	s.db, err = core.Connect(c.GetData())
	if err != nil {
		return nil, err
	}

	return &s, err
}

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   *cfg.Auth
	configDB *cfg.Data
}

func (s Service) Authenticate(credentials []Credential) (ID, error) {
	var id ID
	var err error

	// TODO: Implement Authenticate

	return id, err
}
