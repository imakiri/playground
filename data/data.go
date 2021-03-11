package data

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/imakiri/playground/erres"
	"github.com/imakiri/playground/transport"
	"github.com/jackc/pgx/v4"
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

func Connect(c *transport.Data) (*pgx.Conn, error) {
	var db *pgx.Conn
	var err error

	if c.GetDSN() == "" {
		return nil, erres.E_InvalidArgument
	}

	db, err = pgx.Connect(context.Background(), c.GetDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
