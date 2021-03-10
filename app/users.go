package app

import (
	"context"
	"github.com/imakiri/playground/core"
)

func (e *User) CreateUser(ctx context.Context, r *core.ContentRequestCreateUser) (core.ContentResponseCreateUser, error) {
	panic("implement me")
}

func (e *User) GetUserInfo(ctx context.Context, r *core.ContentRequestGetUserInfo) (core.ContentResponseGetUserInfo, error) {
	panic("implement me")
}

func (e *User) UpdateUserInfo(ctx context.Context, r *core.ContentRequestUpdateUserInfo) (core.ContentResponseUpdateUserInfo, error) {
	panic("implement me")
}

func (e *User) GetUserList(ctx context.Context, r *core.ContentRequestGetUserList) (core.ContentResponseGetUserList, error) {
	panic("implement me")
}
