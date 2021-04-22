package data

import (
	"github.com/imakiri/gorum/internal/postgres"
)

type serviceUser struct {
	postgres.Connection
}

func (s serviceUser) GetProfile(uuid ModelUserUUID, container *ViewUserProfile) error {
	return postgres.UserGetProfile(s.Connection, uuid, container)
}
func (s serviceUser) UpdateProfile(uuid ModelUserUUID, container ViewUserProfileUpdate) error {
	return postgres.UserUpdateProfile(s.Connection, uuid, container)
}
