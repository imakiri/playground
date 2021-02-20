package misc

import (
	"context"
	"github.com/imakiri/playground/core"
)

// }
// }

// Request {

type DataRequestCreateUser struct {
	core.Meta
	Options     DataViewUserOptions
	Credentials DataViewCredentialsGeneral
}
type DataRequestGetUserInfo struct {
	core.Meta
	ID DataModelUserID
}
type DataRequestUpdateUserInfo struct {
	core.Meta
	ID          DataModelUserID
	Options     DataViewUserOptions
	Credentials DataViewCredentialsGeneral
}
type DataRequestGetUserList struct {
	core.Meta
	Page       uint8
	TotalPages uint8
}
type DataRequestGetThreadContent struct {
	core.Meta
	Page       uint8
	TotalPages uint8
}
type DataRequestPostToThread struct {
	core.Meta
	Post DataModelPostContent
}
type DataRequestGetUserPassHash struct {
	core.Meta
	Login DataModelCredentialsLogin
}
type DataRequestGetUserPermissions struct {
	core.Meta
	ID DataModelUserID
}
type DataRequestGetPostInfo struct {
	core.Meta
	PostUUID DataModelPostUUID
}

// }

// Response {

type DataResponseCreateUser struct {
	core.Meta
}
type DataResponseGetUserInfo struct {
	core.Meta
	Info DataViewUserPublicInfoExt
}
type DataResponseUpdateUserInfo struct {
	core.Meta
}
type DataResponseGetUserList struct {
	core.Meta
	Users []DataViewUserPublicInfo
}
type DataResponseGetThreadContent struct {
	core.Meta
	Posts []DataViewPostDetails
}
type DataResponsePostToThread struct {
	core.Meta
}
type DataResponseGetUserPassHash struct {
	core.Meta
	PassHash DataModelCredentialsPassHash
}
type DataResponseGetUserPermissions struct {
	core.Meta
	Permissions DataModelCredentialsPermissions
}
type DataResponseGetPostInfo struct {
	core.Meta
	Info DataViewPostDetails
}

// }

type DataService interface {
	CreateUser(ctx context.Context, r *DataRequestCreateUser) *DataResponseCreateUser
	GetUserInfo(ctx context.Context, r *DataRequestGetUserInfo) *DataResponseGetUserInfo
	UpdateUserInfo(ctx context.Context, r *DataRequestUpdateUserInfo) *DataResponseUpdateUserInfo
	GetUserList(ctx context.Context, r *DataRequestGetUserList) *DataResponseGetUserList
	GetThreadContent(ctx context.Context, r *DataRequestGetThreadContent) *DataResponseGetThreadContent
	GetPostInfo(ctx context.Context, r *DataRequestGetPostInfo) *DataResponseGetPostInfo
	PostToThread(ctx context.Context, r *DataRequestPostToThread) *DataResponsePostToThread
	GetUserPassHash(ctx context.Context, r *DataRequestGetUserPassHash) *DataResponseGetUserPassHash
	GetUserPermissions(ctx context.Context, r *DataRequestGetUserPermissions) *DataResponseGetUserPermissions
}

// }
