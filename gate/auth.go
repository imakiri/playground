package gate

import "context"

func (e *Service) Login(ctx context.Context, r *RequestLogin) (*ResponseLogin, error) {
	panic("implement me")
}

func (e *Service) Logout(ctx context.Context, r *RequestLogout) error {
	panic("implement me")
}
