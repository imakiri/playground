package data

import (
	"github.com/imakiri/gorum/internal/postgres"
)

type serviceCookie struct {
	postgres.Connection
}

func (s serviceCookie) Add(cookie ModelCookie) error {
	return postgres.CookieAdd(s.Connection, cookie)
}
func (s serviceCookie) Get(key ModelCookieKey, container *ViewCookie) error {
	return postgres.CookieGet(s.Connection, key, container)
}
func (s serviceCookie) Delete(uuid ModelUserUUID) error {
	return postgres.CookieDelete(s.Connection, uuid)
}
