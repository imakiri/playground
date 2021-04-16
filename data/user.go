package data

import (
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
)

type User interface {
	GetProfile(uuid types.ModelUserUUID, container *types.ViewUserProfile) error
	UpdateProfile(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate) error
}

type user struct {
	*service
}

func (s user) GetProfile(uuid types.ModelUserUUID, container *types.ViewUserProfile) error {
	return postgres.UserGetProfile(uuid, container, s.db)
}
func (s user) UpdateProfile(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate) error {
	return postgres.UserUpdateProfile(uuid, container, s.db)
}
