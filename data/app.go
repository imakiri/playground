package data

import (
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
)

func (s Service) GetUserProfile(uuid types.ModelUserUUID, container *types.ViewUserProfile) error {
	return postgres.GetUserProfile(uuid, container, s.db)
}
func (s Service) UpdateUserProfile(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate) error {
	return postgres.UpdateUserProfile(uuid, container, s.db)
}
func (s Service) CreateThread(container types.ViewThreadCreate) error {
	return postgres.CreateThread(container, s.db)
}
func (s Service) GetThread(thread_uuid types.ModelThreadUUID, container *types.ViewThread) error {
	return postgres.GetThread(thread_uuid, container, s.db)
}
func (s Service) GetThreads(category types.ModelCategoryUUID, container *types.ViewThreadsByCategory) error {
	return postgres.GetThreads(category, container, s.db)
}
func (s Service) UpdateThread(uuid types.ModelThreadUUID, container types.ViewThreadUpdate) error {
	var err error

	//

	return err
}
func (s Service) DeleteThread(uuid types.ModelThreadUUID) error {
	var err error

	//

	return err
}
func (s Service) CreatePost(container types.ViewPostCreate) error {
	var err error

	//

	return err
}
func (s Service) UpdatePost(uuid types.ModelPostUUID, container types.ViewPostCreate) error {
	var err error

	//

	return err
}
func (s Service) DeletePost(uuid types.ModelPostUUID) error {
	var err error

	//

	return err
}
