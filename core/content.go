package core

import (
	"context"
	"github.com/imakiri/playground/data"
)

// Content {

// Request {

type ContentRequestGetThreadContent struct {
	Page       uint8
	TotalPages uint8
}
type ContentRequestPostToThread struct {
	data.ModelPostContent
}
type ContentRequestCreateUser struct {
	data.ViewUserOptions
}
type ContentRequestGetUserInfo struct {
	data.ModelUserID
}
type ContentRequestUpdateUserInfo struct {
	data.ModelUserID
	data.ViewUserOptions
}
type ContentRequestGetUserList struct {
	Page       uint8
	TotalPages uint8
}

// }

// Response {

type ContentResponseGetThreadContent struct {
	Meta
	Posts []data.ViewPostDetails
}
type ContentResponsePostToThread struct {
	Meta
}
type ContentResponseCreateUser struct {
	Meta
}
type ContentResponseGetUserInfo struct {
	Meta
	data.ViewUserPublicInfoExt
}
type ContentResponseUpdateUserInfo struct {
	Meta
}
type ContentResponseGetUserList struct {
	Meta
	Users []data.ViewUserPublicInfo
}
type ContentResponseGetTrace struct {
	Meta
	RequestedTrace Meta
}

// }

type ContentService interface {
	GetThreadContent(ctx context.Context, r *ContentRequestGetThreadContent) (*ContentResponseGetThreadContent, error)
	PostToThread(ctx context.Context, r *ContentRequestPostToThread) (*ContentResponsePostToThread, error)
	CreateUser(ctx context.Context, r *ContentRequestCreateUser) (ContentResponseCreateUser, error)
	GetUserInfo(ctx context.Context, r *ContentRequestGetUserInfo) (ContentResponseGetUserInfo, error)
	UpdateUserInfo(ctx context.Context, r *ContentRequestUpdateUserInfo) (ContentResponseUpdateUserInfo, error)
	GetUserList(ctx context.Context, r *ContentRequestGetUserList) (ContentResponseGetUserList, error)
}

// }
