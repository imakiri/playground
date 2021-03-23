package data

import (
	"context"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
)

type (
	ViewUserID struct {
		UUID  ModelUserUUID
		PemID ModelUserPemID
	}
	ViewUserProfile struct {
		Nickname         ModelUserNickname
		Fullname         ModelUserFullname
		Avatar           ModelUserAvatar
		RegistrationDate ModelUserRegistrationDate
	}
	ViewUserProfile4Update struct {
		Nickname *ModelUserNickname
		Fullname *ModelUserFullname
		Avatar   *ModelUserAvatar
	}
	ViewPostHeader struct {
		UserUUID ModelUserUUID
		Content  ModelPostContent
	}
	ViewPostCreate struct {
		UserUUID   ModelUserUUID
		ThreadUUID ModelThreadUUID
		Content    ModelPostContent
	}
	ViewPostUpdate struct {
		UUID    ModelPostUUID
		Content ModelPostContent
	}
	ViewPostByThreadUUID struct {
		UserUUID ModelUserUUID
		PostUUID ModelPostUUID
		Date     ModelPostDate
		LastEdit ModelPostLastEdit
		Content  ModelPostContent
	}
	ViewThreadID struct {
		UUID  ModelThreadUUID
		PemID ModelThreadPemID
	}
	ViewThreadCreate struct {
		UserUUID     ModelUserUUID
		PemID        ModelThreadPemID
		CreationDate ModelThreadCreationDate
		Category     ModelThreadCategory
		Name         ModelThreadName
		Header       ViewPostHeader
	}
	ViewThreadContent struct {
		CreationDate ModelThreadCreationDate
		Category     ModelThreadCategory
		Name         ModelThreadName
		Users        []ModelUser
		Posts        []ViewPostByThreadUUID
	}
	ViewThreadUpdate struct {
		PemID    *ModelThreadPemID
		Category *ModelThreadCategory
		Name     *ModelThreadName
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

// ServiceUser.GetProfile()
// ServiceUser.UpdateProfile()
// ServiceThread.Create()
// ServiceThread.GetThreadContent()
// ServiceThread.GetThreadsList()
// ServiceThread.Update()
// ServiceThread.Delete()
// ServicePost.Create()
// ServicePost.Update()
// ServicePost.Delete()

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
