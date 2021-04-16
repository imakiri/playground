package data

import (
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
)

type Post interface {
	Create(container types.ViewPostCreate) error
	Update(uuid types.ModelPostUUID, container types.ViewPostCreate) error
	Delete(uuid types.ModelPostUUID) error
}

type post struct {
	*service
}

func (s post) Create(container types.ViewPostCreate) error {
	return postgres.PostCreate(container, s.db)
}
func (s post) Update(uuid types.ModelPostUUID, container types.ViewPostCreate) error {
	return postgres.PostUpdate(uuid, container, s.db)
}
func (s post) Delete(uuid types.ModelPostUUID) error {
	return postgres.PostDelete(uuid, s.db)
}
