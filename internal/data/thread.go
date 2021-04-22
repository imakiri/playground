package data

import (
	"github.com/imakiri/gorum/internal/postgres"
)

type serviceThread struct {
	postgres.Connection
}

func (s serviceThread) Create(container ModelThread) error {
	return postgres.ThreadCreate(s.Connection, container)
}
func (s serviceThread) Get(thread_uuid ModelThreadUUID, container *ViewThread) error {
	return postgres.ThreadGet(s.Connection, thread_uuid, container)
}
func (s serviceThread) GetList(category ModelCategoryUUID, container *ViewThreadsByCategory) error {
	return postgres.ThreadGetList(s.Connection, category, container)
}
func (s serviceThread) Update(uuid ModelThreadUUID, container ViewThreadUpdate) error {
	return postgres.ThreadUpdate(s.Connection, uuid, container)
}
func (s serviceThread) Delete(uuid ModelThreadUUID) error {
	return postgres.ThreadDelete(s.Connection, uuid)
}
