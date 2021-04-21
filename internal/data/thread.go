package data

import (
	"github.com/imakiri/gorum/internal/postgres"
	"github.com/imakiri/gorum/internal/types"
)

type Thread interface {
	Create(container types.ModelThread) error
	Get(thread_uuid types.ModelThreadUUID, container *types.ViewThread) error
	GetList(category types.ModelCategoryUUID, container *types.ViewThreadsByCategory) error
	Update(uuid types.ModelThreadUUID, container types.ViewThreadUpdate) error
	Delete(uuid types.ModelThreadUUID) error
}

type thread struct {
	*service
}

func (s thread) Create(container types.ModelThread) error {
	return postgres.ThreadCreate(container, s.db)
}
func (s thread) Get(thread_uuid types.ModelThreadUUID, container *types.ViewThread) error {
	return postgres.ThreadGet(thread_uuid, container, s.db)
}
func (s thread) GetList(category types.ModelCategoryUUID, container *types.ViewThreadsByCategory) error {
	return postgres.ThreadGetList(category, container, s.db)
}
func (s thread) Update(uuid types.ModelThreadUUID, container types.ViewThreadUpdate) error {
	return postgres.ThreadUpdate(uuid, container, s.db)
}
func (s thread) Delete(uuid types.ModelThreadUUID) error {
	return postgres.ThreadDelete(uuid, s.db)
}
