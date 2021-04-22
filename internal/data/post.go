package data

import (
	"github.com/imakiri/gorum/internal/postgres"
)

type servicePost struct {
	postgres.Connection
}

func (s servicePost) Create(container ViewPostCreate) error {
	return postgres.PostCreate(s.Connection, container)
}
func (s servicePost) Update(uuid ModelPostUUID, container ViewPostCreate) error {
	return postgres.PostUpdate(s.Connection, uuid, container)
}
func (s servicePost) Delete(uuid ModelPostUUID) error {
	return postgres.PostDelete(s.Connection, uuid)
}
