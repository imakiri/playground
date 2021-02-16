package core

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Data {
// Model {
// User {

type DataModelUserID uint64
type DataModelUserName string
type DataModelUserAvatar []byte
type DataModelUserEmail string
type DataModelUserJoinDate timestamp.Timestamp

// }

// Credentials {

type DataModelCredentialsUserID DataModelUserID
type DataModelCredentialsLogin string
type DataModelCredentialsPassHash []byte
type DataModelCredentialsPermissions []byte

//

// Post {

type DataModelPostUserID DataModelUserID
type DataModelPostUUID uint64
type DataModelPostTimestamp timestamp.Timestamp
type DataModelPostContent string

// }
// }

// View {
// User {

type DataViewUserGeneral struct {
	DataModelUserID
	DataModelUserName
	DataModelUserAvatar
	DataModelUserEmail
	DataModelUserJoinDate
}
type DataViewUserStats struct {
	TotalPosts uint64
}
type DataViewUserPublicInfo struct {
	DataModelUserID
	DataModelUserName
	DataModelUserAvatar
}
type DataViewUserPublicInfoExt struct {
	DataViewUserPublicInfo
	DataModelUserJoinDate
	DataViewUserStats
}
type DataViewUserOptions struct {
	DataModelUserName
	DataModelUserAvatar
	DataModelUserEmail
}

// }

// Credentials {

type DataViewCredentialsGeneral struct {
	DataModelCredentialsLogin
	DataModelCredentialsPassHash
	DataModelCredentialsPermissions
}

// Post {

type DataViewPostGeneral struct {
	DataModelPostUserID
	DataModelPostUUID
	DataModelPostTimestamp
	DataModelPostContent
}
type DataViewPostDetails struct {
	DataModelPostUUID
	DataModelPostTimestamp
	DataModelPostContent
	DataViewUserPublicInfoExt
}

// }
// }

// Request {

type DataRequestCreateUser struct {
	Trace
	Options     DataViewUserOptions
	Credentials DataViewCredentialsGeneral
}
type DataRequestGetUserInfo struct {
	Trace
	ID DataModelUserID
}
type DataRequestUpdateUserInfo struct {
	Trace
	ID          DataModelUserID
	Options     DataViewUserOptions
	Credentials DataViewCredentialsGeneral
}
type DataRequestGetUserList struct {
	Trace
	Page       uint8
	TotalPages uint8
}
type DataRequestGetThreadContent struct {
	Trace
	Page       uint8
	TotalPages uint8
}
type DataRequestPostToThread struct {
	Trace
	Post DataModelPostContent
}
type DataRequestGetUserPassHash struct {
	Trace
	Login DataModelCredentialsLogin
}
type DataRequestGetUserPermissions struct {
	Trace
	ID DataModelUserID
}
type DataRequestGetPostInfo struct {
	Trace
	PostUUID DataModelPostUUID
}

// }

// Response {

type DataResponseCreateUser struct {
	Trace
}
type DataResponseGetUserInfo struct {
	Trace
	Info DataViewUserPublicInfoExt
}
type DataResponseUpdateUserInfo struct {
	Trace
}
type DataResponseGetUserList struct {
	Trace
	Users []DataViewUserPublicInfo
}
type DataResponseGetThreadContent struct {
	Trace
	Posts []DataViewPostDetails
}
type DataResponsePostToThread struct {
	Trace
}
type DataResponseGetUserPassHash struct {
	Trace
	PassHash DataModelCredentialsPassHash
}
type DataResponseGetUserPermissions struct {
	Trace
	Permissions DataModelCredentialsPermissions
}
type DataResponseGetPostInfo struct {
	Trace
	Info DataViewPostDetails
}

// }

type DataService interface {
	CreateUser(DataRequestCreateUser) DataResponseCreateUser
	GetUserInfo(DataRequestGetUserInfo) DataResponseGetUserInfo
	UpdateUserInfo(DataRequestUpdateUserInfo) DataResponseUpdateUserInfo
	GetUserList(DataRequestGetUserList) DataResponseGetUserList
	GetThreadContent(DataRequestGetThreadContent) DataResponseGetThreadContent
	PostToThread(DataRequestPostToThread) DataResponsePostToThread
	GetUserPassHash(DataRequestGetUserPassHash) DataResponseGetUserPassHash
	GetUserPermissions(DataRequestGetUserPermissions) DataResponseGetUserPermissions
	GetPostInfo(DataRequestGetPostInfo) DataResponseGetPostInfo
}

// }
