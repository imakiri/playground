package data

import (
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
)

func (s Service) AddCookie(uuid types.ModelUserUUID, cookie types.ViewCookieByUUID) error {
	return postgres.AddCookieV1(uuid, cookie, s.db)
}

func (s Service) GetCookie(key types.ModelCookieKey, container *types.ViewCookieByUUID) error {
	return postgres.GetCookieV1(key, container, s.db)
}

func (s Service) DeleteCookie(uuid types.ModelUserUUID) error {
	return postgres.DeleteCookieV1(uuid, s.db)
}

func (s Service) AddLogpass(uuid types.ModelUserUUID, logpass types.ViewLogpassByUUID) error {
	return postgres.AddLogpassV1(uuid, logpass, s.db)
}

func (s Service) GetLogpass(login types.ModelLogpassLogin, container *types.ViewLogpassByUUID) error {
	return postgres.GetLogpassV1(login, container, s.db)
}

func (s Service) DeleteLogpass(uuid types.ModelUserUUID) error {
	return postgres.DeleteLogpassV1(uuid, s.db)
}
