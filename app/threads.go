package app

import (
	"context"
	"github.com/imakiri/playground/core"
)

func (e *User) GetThreadContent(ctx context.Context, r *core.ContentRequestGetThreadContent) (*core.ContentResponseGetThreadContent, error) {
	panic("implement me")
}

func (e *User) PostToThread(ctx context.Context, r *core.ContentRequestPostToThread) (*core.ContentResponsePostToThread, error) {
	panic("implement me")
}
