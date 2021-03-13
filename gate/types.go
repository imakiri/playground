package gate

import (
	"context"
	"github.com/imakiri/gorum/core"
	"github.com/imakiri/gorum/data"
)

// Gate

// Request

type RequestGetThreadContent struct {
	Key core.AuthKey
	core.ContentRequestGetThreadContent
}
type RequestPostToThread struct {
	Key core.AuthKey
	core.ContentRequestPostToThread
}

type RequestCreateUser struct {
	Key core.AuthKey
	core.ContentRequestCreateUser
}
type RequestGetUserInfo struct {
	Key core.AuthKey
	core.ContentRequestGetUserInfo
}
type RequestUpdateUserInfo struct {
	Key core.AuthKey
	core.ContentRequestUpdateUserInfo
}
type RequestGetUserList struct {
	Key core.AuthKey
	core.ContentRequestGetUserList
}

type RequestLogin struct {
	core.AuthRequestLogin
}
type RequestLogout struct {
	core.AuthRequestLogout
}

//

// Response

type ResponseGetThreadContent struct {
	Posts []data.ViewPostDetails
}

type ResponseCreateUser struct {
	core.AuthKey
}
type ResponseGetUserInfo struct {
	data.ViewUserPublicInfoExt
}
type ResponseGetUserList struct {
	core.Meta
	Users []data.ViewUserPublicInfo
}

type ResponseLogin struct {
	core.AuthKey
}

type UserService interface {
	RegisterUser(ctx context.Context, r *RequestCreateUser) (*ResponseCreateUser, error)
	GetUserInfo(ctx context.Context, r *RequestGetUserInfo) (*ResponseGetUserInfo, error)
	GetUserList(ctx context.Context, r *RequestGetUserList) (*ResponseGetUserList, error)
	UpdateUserInfo(ctx context.Context, r *RequestUpdateUserInfo) error
}
type ThreadService interface {
	GetThreadContent(ctx context.Context, r *RequestGetThreadContent) (*ResponseGetThreadContent, error)
	PostToThread(ctx context.Context, r *RequestPostToThread) error
}
type AuthService interface {
	Login(ctx context.Context, r *RequestLogin) (*ResponseLogin, error)
	Logout(ctx context.Context, r *RequestLogout) error
}
type GeneralService interface {
	UserService
	ThreadService
	AuthService
}

//
