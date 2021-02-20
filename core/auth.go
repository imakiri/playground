package core

import (
	"context"
	"github.com/imakiri/playground/data"
)

// Auth

type AuthKey []byte

// Request

type AuthRequestLogin struct {
	Login    data.ModelCredentialsLogin
	Password string
}
type AuthRequestCheckAccess struct {
	Key AuthKey
}
type AuthRequestLogout struct {
	Key AuthKey
}

//

// Response

type AuthResponseLogin struct {
	Meta
	Key AuthKey
}
type AuthResponseLogout struct {
	Meta
}

//

type AuthService interface {
	Login(ctx context.Context, r *AuthRequestLogin) (*AuthResponseLogin, error)
	CheckAccess(ctx context.Context, r *AuthRequestCheckAccess) error
	Logout(ctx context.Context, r *AuthRequestLogout) (*AuthResponseLogout, error)
}

//
