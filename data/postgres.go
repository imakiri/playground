package data

import (
	"context"
	"github.com/aidarkhanov/nanoid"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/cfg"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"time"
)

type (
	ViewUserID struct {
		UserUUID ModelUserUUID
		PemID    ModelUserPemID
	}
	ViewUserAvatar struct {
		Avatar512 ModelUserAvatar512
		Avatar256 ModelUserAvatar256
		Avatar128 ModelUserAvatar128
	}
)

type (
	ViewUserProfile struct {
		RegistrationDate ModelUserRegistrationDate
		NickName         ModelUserNickName
		FullName         ModelUserFullName
		Avatar512        ModelUserAvatar512
	}
	ViewUserProfileFromThread struct {
		UserUUID         ModelUserUUID
		RegistrationDate ModelUserRegistrationDate
		NickName         ModelUserNickName
		FullName         ModelUserFullName
		Avatar256        ModelUserAvatar256
	}
	ViewUserProfileFromMain struct {
		UserUUID  ModelUserUUID
		NickName  ModelUserNickName
		Avatar128 ModelUserAvatar128
	}
	ViewUserProfileUpdate struct {
		NickName *ModelUserNickName
		FullName *ModelUserFullName
		Avatar   *ViewUserAvatar
	}
	ViewPostCreate struct {
		UserUUID   ModelUserUUID
		ThreadUUID ModelThreadUUID
		Content    ModelContent
	}
	ViewPostUpdate struct {
		Content ModelContent
	}
	ViewPostByThreadUUID struct {
		PostUUID     ModelPostUUID
		UserUUID     ModelUserUUID
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Content      ModelContent
	}
	ViewThreadCreate struct {
		CategoryUUID ModelCategoryUUID
		UserUUID     ModelUserUUID
		Name         ModelThreadName
		Header       ModelContent
	}
	ViewThread struct {
		Category     ModelCategory
		Author       ViewUserProfileFromThread
		Name         ModelThreadName
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Header       ModelContent
		Content      ViewThreadContent
	}
	ViewThreadContent struct {
		Users []ViewUserProfileFromThread
		Posts []ViewPostByThreadUUID
	}
	ViewThreadsByCategory struct {
		ThreadUUID ModelThreadUUID
		Name       ModelThreadName
	}
	ViewThreadUpdate struct {
		CategoryUUID *ModelCategoryUUID
		Name         *ModelThreadName
		Header       *ModelContent
	}
)

type ConfigApp interface {
	Get4DataApp(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.DataApp, error)
}

type ServicePostgres struct {
	config       ConfigApp
	configCached *cfg.DataApp
	db           *sqlx.DB
}

func (p ServicePostgres) GetUserProfile(uuid ModelUserUUID, container *ViewUserProfile) error {
	var err = p.db.Get(container, "SELECT registration_date, nick_name, full_name, avatar512 FROM main.app.users", uuid)
	return errrapper(err)
}

func (p ServicePostgres) UpdateUserProfile(uuid ModelUserUUID, container ViewUserProfileUpdate) error {
	var tx, err = p.db.Begin()
	if err != nil {
		return errrapper(err)
	}

	if container.Avatar != nil {
		_, err = tx.Exec("UPDATE app.users SET avatar512 = $2, avatar256 = $3, avatar128 = $4 WHERE user_uuid = $1", uuid, container.Avatar.Avatar512, container.Avatar.Avatar256, container.Avatar.Avatar128)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errrapper(e).AddRoute("Avatar")
			}
			return errrapper(err).AddRoute("Avatar")
		}
	}
	if container.FullName != nil {
		_, err = tx.Exec("UPDATE app.users SET full_name = $2 WHERE user_uuid = $1", uuid, container.FullName)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errrapper(e).AddRoute("Fullname")
			}
			return errrapper(err).AddRoute("Fullname")
		}
	}
	if container.NickName != nil {
		_, err = tx.Exec("UPDATE app.users SET nick_name = $2 WHERE user_uuid = $1", uuid, container.NickName)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errrapper(e).AddRoute("Nickname")
			}
			return errrapper(err).AddRoute("Nickname")
		}
	}

	err = tx.Commit()
	return errrapper(err)
}

func (p ServicePostgres) CreateThread(container ViewThreadCreate) error {
	var threadUUID = ModelThreadUUID(nanoid.New())
	var now = ModelDate(time.Now().UnixNano())
	var thread = ModelThread{
		ThreadUUID:   threadUUID,
		CategoryUUID: container.CategoryUUID,
		UserUUID:     container.UserUUID,
		Name:         container.Name,
		DateAdded:    now,
		DateLastEdit: now,
		Header:       container.Header,
	}

	var _, err = p.db.NamedExec("INSERT INTO app.threads VALUES (:ThreadUUID, :CategoryUUID, :UserUUID, :Name, :DateAdded, :DateLastEdit, :Header)", thread)
	return errrapper(err)
}

func (p ServicePostgres) GetThread(thread_uuid ModelThreadUUID, container *ViewThread) error {
	var err error

	err = p.db.Get(container, "SELECT category_uuid, name FROM app.threads WHERE thread_uuid = $1", thread_uuid) // ----------------------
	if err != nil {
		return errrapper(err)
	}

	err = p.db.Get(container.Content.Posts, "SELECT user_uuid, post_uuid, date_added, date_last_edit, content  FROM app.posts WHERE thread_uuid = $1", thread_uuid)
	if err != nil {
		return errrapper(err)
	}

	err = p.db.Get(container.Content.Users, "SELECT DISTINCT users.user_uuid, registration_date, nick_name, full_name, avatar256 FROM app.posts INNER JOIN app.users ON users.user_uuid = posts.user_uuid WHERE thread_uuid = $1", thread_uuid)
	if err != nil {
		return errrapper(err)
	}

	return errrapper(err)
}

func (p ServicePostgres) GetThreads(category ModelCategoryUUID, container *ViewThreadsByCategory) error {
	var err = p.db.Get(container, "", category)
	return errrapper(err)
}

func (p ServicePostgres) UpdateThread(uuid ModelThreadUUID, container ViewThreadUpdate) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) DeleteThread(uuid ModelThreadUUID) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) CreatePost(container ViewPostCreate) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) UpdatePost(uuid ModelPostUUID, container ViewPostCreate) error {
	var err error

	//

	return errrapper(err)
}

func (p ServicePostgres) DeletePost(uuid ModelPostUUID) error {
	var err error

	//

	return errrapper(err)
}

func NewServicePostgres(c ConfigApp) (*ServicePostgres, error) {
	var s ServicePostgres
	var err error

	s.config = c
	s.configCached, err = s.config.Get4DataApp(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	s.db, err = sqlx.Connect("pgx", s.configCached.GetDSN())
	if err != nil {
		return nil, erres.ConnectionError.Extend().AddDescription(err.Error())
	}
	return &s, err
}
