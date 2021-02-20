package content

import (
	"context"
	"github.com/imakiri/playground/core"
)

func (e *Service) CreateUser(ctx context.Context, r *core.ContentRequestCreateUser) (core.ContentResponseCreateUser, error) {
	panic("implement me")
}

func (e *Service) GetUserInfo(ctx context.Context, r *core.ContentRequestGetUserInfo) (core.ContentResponseGetUserInfo, error) {
	panic("implement me")
}

func (e *Service) UpdateUserInfo(ctx context.Context, r *core.ContentRequestUpdateUserInfo) (core.ContentResponseUpdateUserInfo, error) {
	panic("implement me")
}

func (e *Service) GetUserList(ctx context.Context, r *core.ContentRequestGetUserList) (core.ContentResponseGetUserList, error) {
	panic("implement me")
}
