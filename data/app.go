package data

import (
	"context"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
)

type (
	ViewPostByThreadUUID struct {
		UserUUID ModelUserUUID
		PostUUID ModelPostUUID
		Date     ModelPostDate
		LastEdit ModelPostLastEdit
		Content  ModelPostContent
	}
	ViewThread struct {
		CreationDate ModelThreadCreationDate
		Topic        ModelThreadTopic
		Name         ModelThreadName
		Users        []ModelUser
		Posts        []ViewPostByThreadUUID
	}
)

type ConfigApp interface {
	Get4DataApp(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.DataApp, error)
}

type App struct {
	config       ConfigApp
	configCached *cfg.DataApp
	db           *sqlx.DB
}

func NewApp(c ConfigApp) (*App, error) {
	var s App
	var err error

	s.config = c
	s.configCached, err = s.config.Get4DataApp(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	s.db, err = sqlx.Connect("pgx", s.configCached.GetDSN())
	if err != nil {
		return nil, erres.ConnectionError.SetTime("").SetDescription(err.Error())
	}
	return &s, err
}
