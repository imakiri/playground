package data

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type (
	ModelUserUUID             string
	ModelUserPemID            int16
	ModelUserNickname         string
	ModelUserFullname         string
	ModelUserAvatar           []byte
	ModelUserRegistrationDate int64

	UserID struct {
		UUID  ModelUserUUID
		PemID ModelUserPemID
	}
)

type ModelPostUserID UserID
type ModelPostUUID uint64
type ModelPostTimestamp timestamp.Timestamp
type ModelPostContent string

type ViewUserGeneral struct {
	ID       UserID
	Name     ModelUserNickname
	Avatar   ModelUserAvatar
	JoinDate ModelUserRegistrationDate
}
type ViewUserStats struct {
	TotalPosts uint64
}
type ViewUserPublicInfo struct {
	ID     UserID
	Name   ModelUserNickname
	Avatar ModelUserAvatar
}
type ViewUserPublicInfoExt struct {
	ViewUserPublicInfo
	JoinDate ModelUserRegistrationDate
	ViewUserStats
}
type ViewUserOptions struct {
	Name   ModelUserNickname
	Avatar ModelUserAvatar
}

type ViewPostGeneral struct {
	UUID      ModelPostUUID
	UserID    ModelPostUserID
	Timestamp ModelPostTimestamp
	Content   ModelPostContent
}
type ViewPostDetails struct {
	UUID      ModelPostUUID
	Timestamp ModelPostTimestamp
	Content   ModelPostContent
	ViewUserPublicInfoExt
}
