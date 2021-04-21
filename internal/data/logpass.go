package data

import (
	"github.com/imakiri/gorum/internal/postgres"
	"github.com/imakiri/gorum/internal/types"
)

type Logpass interface {
	Add(logpass types.ModelLogpass) error
	Get(login types.ModelLogpassLogin, container *types.ViewLogpass) error
	Delete(uuid types.ModelUserUUID) error
}

type logpass struct {
	*service
}

func (s logpass) Add(logpass types.ModelLogpass) error {
	return postgres.LogpassAdd(logpass, s.db)
}
func (s logpass) Get(login types.ModelLogpassLogin, container *types.ViewLogpass) error {
	return postgres.LogpassGetWithLogin(login, container, s.db)
}
func (s logpass) Delete(uuid types.ModelUserUUID) error {
	return postgres.LogpassDelete(uuid, s.db)
}
