package core

// App {

// Request {

type AppRequestGetThreadContent struct {
	Page       uint8
	TotalPages uint8
}
type AppRequestPostToThread struct {
	DataModelPostContent
}
type AppRequestCreateUser struct {
	DataViewUserOptions
}
type AppRequestGetUserInfo struct {
	DataModelUserID
}
type AppRequestUpdateUserInfo struct {
	DataModelUserID
	DataViewUserOptions
}
type AppRequestGetUserList struct {
	Page       uint8
	TotalPages uint8
}
type AppRequestGetTrace struct {
	AuthAccessKey
	ActionID
}

// }

// Response {

type AppResponseGetThreadContent struct {
	Trace
	Posts []DataViewPostDetails
}
type AppResponsePostToThread struct {
	Trace
}
type AppResponseCreateUser struct {
	Trace
}
type AppResponseGetUserInfo struct {
	Trace
	DataViewUserPublicInfoExt
}
type AppResponseUpdateUserInfo struct {
	Trace
}
type AppResponseGetUserList struct {
	Trace
	Users []DataViewUserPublicInfo
}
type AppResponseGetTrace struct {
	Trace
	RequestedTrace Trace
}

// }

type AppService interface {
	GetThreadContent(AppRequestGetThreadContent) AppResponseGetThreadContent
	PostToThread(AppRequestPostToThread) AppResponsePostToThread
	CreateUser(AppRequestCreateUser) AppResponseCreateUser
	GetUserInfo(AppRequestGetUserInfo) AppResponseGetUserInfo
	UpdateUserInfo(AppRequestUpdateUserInfo) AppResponseUpdateUserInfo
	GetUserList(AppRequestGetUserList) AppResponseGetUserList
	GetTrace(AppRequestGetTrace) AppResponseGetTrace
}

// }
