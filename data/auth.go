package data

import (
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
)

func (s Service) AddCookie(uuid types.ModelUserUUID, cookie types.ViewCookieByUUID) error {
	return postgres.AddCookie(uuid, cookie, s.db)
}

func (s Service) GetCookie(key types.ModelCookieKey, container *types.ViewCookieByUUID) error {
	return postgres.GetCookie(key, container, s.db)
}

func (s Service) DeleteCookie(uuid types.ModelUserUUID) error {
	return postgres.DeleteCookie(uuid, s.db)
}

func (s Service) AddLogpass(uuid types.ModelUserUUID, logpass types.ViewLogpassByUUID) error {
	return postgres.AddLogpass(uuid, logpass, s.db)
}

func (s Service) GetLogpass(login types.ModelLogpassLogin, container *types.ViewLogpassByUUID) error {
	return postgres.GetLogpass(login, container, s.db)
}

func (s Service) DeleteLogpass(uuid types.ModelUserUUID) error {
	return postgres.DeleteLogpass(uuid, s.db)
}
