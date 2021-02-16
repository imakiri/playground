package core

// Gate

// Request

type GateRequestGetThreadContent struct {
	AuthAccessKey
	AppRequestGetThreadContent
}
type GateRequestPostToThread struct {
	AuthAccessKey
	AppRequestPostToThread
}
type GateRequestCreateUser struct {
	AuthAccessKey
	AppRequestCreateUser
}
type GateRequestGetUserInfo struct {
	AuthAccessKey
	AppRequestGetUserInfo
}
type GateRequestUpdateUserInfo struct {
	AuthAccessKey
	AppRequestUpdateUserInfo
}
type GateRequestGetUserList struct {
	AuthAccessKey
	AppRequestGetUserList
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
	Resume
	Posts []DataViewPostDetails
}
type GateResponsePostToThread struct {
	Resume
}
type GateResponseCreateUser struct {
	Resume
}
type GateResponseGetUserInfo struct {
	Resume
	DataViewUserPublicInfoExt
}
type GateResponseUpdateUserInfo struct {
	Resume
}
type GateResponseGetUserList struct {
	Resume
	Users []DataViewUserPublicInfo
}
type GateResponseLogin struct {
	Resume
	AuthAccessKey
}
type GateResponseLogout struct {
	Resume
}
type GateResponseGetTrace struct {
	Resume
	Trace
}

//

type GateService interface {
	GetThreadContent(GateRequestGetThreadContent) GateResponseGetThreadContent
	PostToThread(GateRequestPostToThread) GateResponsePostToThread
	CreateUser(GateRequestCreateUser) GateResponseCreateUser
	GetUserInfo(GateRequestGetUserInfo) GateResponseGetUserInfo
	UpdateUserInfo(GateRequestUpdateUserInfo) GateResponseUpdateUserInfo
	GetUserList(GateRequestGetUserList) GateResponseGetUserList
	Login(GateRequestLogin) GateResponseLogin
	Logout(GateRequestLogout) GateResponseLogout
	GetReflection(AppRequestGetTrace) GateResponseGetTrace
}

//
