package auth

import (
	"context"
	"github.com/imakiri/playground/core"
)

func (e *Service) Login(ctx context.Context, r *core.AuthRequestLogin) (*core.AuthResponseLogin, error) {
	var re core.AuthResponseLogin
	var err error

	return &re, err
}

func (e *Service) CheckAccess(ctx context.Context, r *core.AuthRequestCheckAccess) error {
	var err error

	return err
}

func (e *Service) Logout(ctx context.Context, r *core.AuthRequestLogout) (*core.AuthResponseLogout, error) {
	panic("implement me")
}
