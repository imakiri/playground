package core

import (
	"context"
	"github.com/imakiri/playground/data"
)

// Gate

// Request

type GateRequestGetThreadContent struct {
	Key AuthKey
	ContentRequestGetThreadContent
}
type GateRequestPostToThread struct {
	Key AuthKey
	ContentRequestPostToThread
}
type GateRequestCreateUser struct {
	Key AuthKey
	ContentRequestCreateUser
}
type GateRequestGetUserInfo struct {
	Key AuthKey
	ContentRequestGetUserInfo
}
type GateRequestUpdateUserInfo struct {
	Key AuthKey
	ContentRequestUpdateUserInfo
}
type GateRequestGetUserList struct {
	Key AuthKey
	ContentRequestGetUserList
}
type GateRequestLogin struct {
	AuthRequestLogin
}
type GateRequestLogout struct {
	AuthRequestLogout
}

//

// Response

type GateResponseGetThreadContent struct {
	Meta
	Posts []data.ViewPostDetails
}
type GateResponsePostToThread struct {
	Meta
}
type GateResponseCreateUser struct {
	Meta
}
type GateResponseGetUserInfo struct {
	Meta
	data.ViewUserPublicInfoExt
}
type GateResponseUpdateUserInfo struct {
	Meta
}
type GateResponseGetUserList struct {
	Meta
	Users []data.ViewUserPublicInfo
}
type GateResponseLogin struct {
	Meta
	AuthKey
}
type GateResponseLogout struct {
	Meta
}

type GateService interface {
	GetThreadContent(ctx context.Context, r *GateRequestGetThreadContent) (*GateResponseGetThreadContent, error)
	PostToThread(ctx context.Context, r *GateRequestPostToThread) (*GateResponsePostToThread, error)
	CreateUser(ctx context.Context, r *GateRequestCreateUser) (*GateResponseCreateUser, error)
	GetUserInfo(ctx context.Context, r *GateRequestGetUserInfo) (*GateResponseGetUserInfo, error)
	UpdateUserInfo(ctx context.Context, r *GateRequestUpdateUserInfo) (*GateResponseUpdateUserInfo, error)
	GetUserList(ctx context.Context, r *GateRequestGetUserList) (*GateResponseGetUserList, error)
	Login(ctx context.Context, r *GateRequestLogin) (*GateResponseLogin, error)
	Logout(ctx context.Context, r *GateRequestLogout) (*GateResponseLogout, error)
}

//
