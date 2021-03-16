package gate

import (
	"context"
	"github.com/imakiri/gorum/data"
	"github.com/imakiri/gorum/utils"
)

// Gate

// Request

type RequestGetThreadContent struct {
	Key utils.AuthKey
	utils.ContentRequestGetThreadContent
}
type RequestPostToThread struct {
	Key utils.AuthKey
	utils.ContentRequestPostToThread
}

type RequestCreateUser struct {
	Key utils.AuthKey
	utils.ContentRequestCreateUser
}
type RequestGetUserInfo struct {
	Key utils.AuthKey
	utils.ContentRequestGetUserInfo
}
type RequestUpdateUserInfo struct {
	Key utils.AuthKey
	utils.ContentRequestUpdateUserInfo
}
type RequestGetUserList struct {
	Key utils.AuthKey
	utils.ContentRequestGetUserList
}

type RequestLogin struct {
	utils.AuthRequestLogin
}
type RequestLogout struct {
	utils.AuthRequestLogout
}

//

// Response

type ResponseGetThreadContent struct {
	Posts []data.ViewPostDetails
}

type ResponseCreateUser struct {
	utils.AuthKey
}
type ResponseGetUserInfo struct {
	data.ViewUserPublicInfoExt
}
type ResponseGetUserList struct {
	utils.Meta
	Users []data.ViewUserPublicInfo
}

type ResponseLogin struct {
	utils.AuthKey
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
