package data

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type ModelUserID uint64
type ModelUserName string
type ModelUserAvatar []byte
type ModelUserEmail string
type ModelUserJoinDate timestamp.Timestamp

type ModelCredentialsUserID ModelUserID
type ModelCredentialsLogin string
type ModelCredentialsPassHash []byte
type ModelCredentialsPermissions []byte

type ModelPostUserID ModelUserID
type ModelPostUUID uint64
type ModelPostTimestamp timestamp.Timestamp
type ModelPostContent string

type ViewUserGeneral struct {
	ID       ModelUserID
	Name     ModelUserName
	Avatar   ModelUserAvatar
	Email    ModelUserEmail
	JoinDate ModelUserJoinDate
}
type ViewUserStats struct {
	TotalPosts uint64
}
type ViewUserPublicInfo struct {
	ID     ModelUserID
	Name   ModelUserName
	Avatar ModelUserAvatar
}
type ViewUserPublicInfoExt struct {
	ViewUserPublicInfo
	JoinDate ModelUserJoinDate
	ViewUserStats
}
type ViewUserOptions struct {
	Name   ModelUserName
	Avatar ModelUserAvatar
	Email  ModelUserEmail
}

type ViewCredentialsGeneral struct {
	Login       ModelCredentialsLogin
	PassHash    ModelCredentialsPassHash
	Permissions ModelCredentialsPermissions
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
