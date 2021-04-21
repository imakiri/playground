package data

import (
	"github.com/imakiri/gorum/internal/postgres"
	"github.com/imakiri/gorum/internal/types"
)

type Cookie interface {
	Add(cookie types.ModelCookie) error
	Get(key types.ModelCookieKey, container *types.ViewCookie) error
	Delete(uuid types.ModelUserUUID) error
}

type cookie struct {
	*service
}

func (s cookie) Add(cookie types.ModelCookie) error {
	return postgres.CookieAdd(cookie, s.db)
}
func (s cookie) Get(key types.ModelCookieKey, container *types.ViewCookie) error {
	return postgres.CookieGet(key, container, s.db)
}
func (s cookie) Delete(uuid types.ModelUserUUID) error {
	return postgres.CookieDelete(uuid, s.db)
}
