package core

// Auth

type AuthAccessKey struct {
	IsWorth bool
	Key     []byte
}

// Request

type AuthRequestLogin struct {
	Trace
	Login    DataModelCredentialsLogin
	Password string
}
type AuthRequestLogout struct {
	Trace
	AuthAccessKey
}

//

// Response

type AuthResponseLogin struct {
	Trace
	AuthAccessKey
}
type AuthResponseLogout struct {
	Trace
}

//

type AuthService interface {
	Login(AuthRequestLogin) AuthResponseLogin
	Logout(AuthRequestLogout) AuthResponseLogout
}

//
