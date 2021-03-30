package data

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/types"
	"github.com/jmoiron/sqlx"
)

type ServicePostgres struct {
	config *types.ConfigDataPostgres
	db     *sqlx.DB
}

// --

func (p ServicePostgres) AddCookie(uuid types.ModelUserUUID, cookie types.ViewCookieByUUID) error {
	return postgres_AddCookie_v1(uuid, cookie, p)
}

func (p ServicePostgres) GetCookie(key types.ModelCookieKey, container *types.ViewCookieByUUID) error {
	return postgres_GetCookie_v1(key, container, p)
}

func (p ServicePostgres) DeleteCookie(uuid types.ModelUserUUID) error {
	return postgres_DeleteCookie_v1(uuid, p)
}

func (p ServicePostgres) AddLogpass(uuid types.ModelUserUUID, logpass types.ViewLogpassByUUID) error {
	return postgres_AddLogpass_v1(uuid, logpass, p)
}

func (p ServicePostgres) GetLogpass(login types.ModelLogpassLogin, container *types.ViewLogpassByUUID) error {
	return postgres_GetLogpass_v1(login, container, p)
}

func (p ServicePostgres) DeleteLogpass(uuid types.ModelUserUUID) error {
	return postgres_DeleteLogpass_v1(uuid, p)
}

// --

func (p ServicePostgres) GetUserProfile(uuid types.ModelUserUUID, container *types.ViewUserProfile) error {
	return postgres_GetUserProfile_v1(uuid, container, p)
}

func (p ServicePostgres) UpdateUserProfile(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate) error {
	return postgres_UpdateUserProfile_v1(uuid, container, p)
}

func (p ServicePostgres) CreateThread(container types.ViewThreadCreate) error {
	return postgres_CreateThread_v1(container, p)
}

func (p ServicePostgres) GetThread(thread_uuid types.ModelThreadUUID, container *types.ViewThread) error {
	return postgres_GetThread_v1(thread_uuid, container, p)
}

func (p ServicePostgres) GetThreads(category types.ModelCategoryUUID, container *types.ViewThreadsByCategory) error {
	return postgres_GetThreads_v1(category, container, p)
}

func (p ServicePostgres) UpdateThread(uuid types.ModelThreadUUID, container types.ViewThreadUpdate) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) DeleteThread(uuid types.ModelThreadUUID) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) CreatePost(container types.ViewPostCreate) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) UpdatePost(uuid types.ModelPostUUID, container types.ViewPostCreate) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) DeletePost(uuid types.ModelPostUUID) error {
	var err error

	//

	return errrapper(err)
}

// --

func NewServicePostgres(c *types.ConfigDataPostgres) (*ServicePostgres, error) {
	var s ServicePostgres
	var err error

	s.config = c

	s.db, err = sqlx.Connect("pgx", s.config.GetDSN())
	if err != nil {
		return nil, erres.ConnectionError.Extend().AddDescription(err.Error())
	}
	return &s, err
}
