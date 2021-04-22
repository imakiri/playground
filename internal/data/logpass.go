package data

import (
	"github.com/imakiri/gorum/internal/postgres"
)

type serviceLogpass struct {
	postgres.Connection
}

func (s serviceLogpass) Add(logpass ModelLogpass) error {
	return postgres.LogpassAdd(s.Connection, logpass)
}
func (s serviceLogpass) Get(login ModelLogpassLogin, container *ViewLogpass) error {
	return postgres.LogpassGetWithLogin(s.Connection, login, container)
}
func (s serviceLogpass) Delete(uuid ModelUserUUID) error {
	return postgres.LogpassDelete(s.Connection, uuid)
}
