package postgres

import (
	"database/sql"
	"github.com/aidarkhanov/nanoid"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/types"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

// Wrapper for raw sql/sqlx/pgx error strings
func errWrapper(err error) *erres.Error {
	switch {
	case err == nil:
		return nil
	case err == sql.ErrTxDone:
		return erres.InternalServiceError.Extend(1).SetDescription(err.Error())
	}

	var e = err.Error()

	switch {
	case strings.Contains(e, "sqlx.bindNamedMapper: unsupported map type:"):
		return erres.InternalServiceError.Extend(1).SetDescription(err.Error())
	default:
		return erres.JustError.Extend(1).SetDescription(err.Error())
	}
}

func AddCookieV1(uuid types.ModelUserUUID, cookie types.ViewCookieByUUID, db *sqlx.DB) error {
	var c types.ModelCookie

	if uuid == "" {
		c.UUID = types.ModelUserUUID(nanoid.New())
	}

	c.Key = cookie.Key
	c.PemID = cookie.PemID
	c.ExpirationDate = cookie.ExpirationDate

	var _, err = db.NamedQuery("INSERT INTO main.auth.cookie VALUES (:key, :uuid, :pemid, :expirationDate)", c)
	return errWrapper(err)
}
func GetCookieV1(key types.ModelCookieKey, container *types.ViewCookieByUUID, db *sqlx.DB) error {
	return db.Get(container, "SELECT uuid, pemid, expiration_date FROM main.auth.cookie WHERE key = $1", key)
}
func DeleteCookieV1(uuid types.ModelUserUUID, db *sqlx.DB) error {
	var _, err = db.Exec("DELETE FROM main.auth.cookie WHERE uuid = $1", uuid)
	return errWrapper(err)
}
func AddLogpassV1(uuid types.ModelUserUUID, logpass types.ViewLogpassByUUID, db *sqlx.DB) error {
	var l types.ModelLogpass

	if uuid == "" {
		l.UUID = types.ModelUserUUID(nanoid.New())
	}

	l.PemID = logpass.PemID
	l.Password = logpass.Password
	l.Login = logpass.Login

	var _, err = db.NamedQuery("INSERT INTO main.auth.logpass VALUES (:uuid, :login, :password, :pemid)", l)
	return errWrapper(err)
}
func GetLogpassV1(login types.ModelLogpassLogin, container *types.ViewLogpassByUUID, db *sqlx.DB) error {
	var err = db.Get(container, "SELECT uuid, pemid, password FROM main.auth.logpass WHERE login = $1", login)
	return errWrapper(err)
}
func DeleteLogpassV1(uuid types.ModelUserUUID, db *sqlx.DB) error {
	var _, err = db.Exec("DELETE FROM main.auth.logpass WHERE uuid = $1", uuid)
	return errWrapper(err)
}

func GetUserProfileV1(uuid types.ModelUserUUID, container *types.ViewUserProfile, db *sqlx.DB) error {
	var err = db.Get(container, "SELECT registration_date, nick_name, full_name, avatar512 FROM main.app.users WHERE user_uuid = $1", uuid)
	return errWrapper(err)
}
func UpdateUserProfileV1(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate, db *sqlx.DB) error {
	var tx, err = db.Begin()
	if err != nil {
		return errWrapper(err)
	}

	if container.Avatar != nil {
		_, err = tx.Exec("UPDATE app.users SET avatar512 = $2, avatar256 = $3, avatar128 = $4 WHERE user_uuid = $1", uuid, container.Avatar.Avatar512, container.Avatar.Avatar256, container.Avatar.Avatar128)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errWrapper(e).SetName("Avatar")
			}
			return errWrapper(err).SetName("Avatar")
		}
	}
	if container.FullName != nil {
		_, err = tx.Exec("UPDATE app.users SET full_name = $2 WHERE user_uuid = $1", uuid, container.FullName)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errWrapper(e).SetName("Fullname")
			}
			return errWrapper(err).SetName("Fullname")
		}
	}
	if container.NickName != nil {
		_, err = tx.Exec("UPDATE app.users SET nick_name = $2 WHERE user_uuid = $1", uuid, container.NickName)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errWrapper(e).SetName("Nickname")
			}
			return errWrapper(err).SetName("Nickname")
		}
	}

	err = tx.Commit()
	return errWrapper(err)
}
func CreateThreadV1(container types.ViewThreadCreate, db *sqlx.DB) error {
	var threadUUID = types.ModelThreadUUID(nanoid.New())
	var now = types.ModelDate(time.Now().UnixNano())
	var thread = types.ModelThread{
		ThreadUUID:   threadUUID,
		CategoryUUID: container.CategoryUUID,
		UserUUID:     container.UserUUID,
		Name:         container.Name,
		DateAdded:    now,
		DateLastEdit: now,
		Header:       container.Header,
	}

	var _, err = db.NamedExec("INSERT INTO app.threads VALUES (:ThreadUUID, :CategoryUUID, :UserUUID, :Name, :DateAdded, :DateLastEdit, :Header)", thread)
	return errWrapper(err)
}
func GetThreadV1(thread_uuid types.ModelThreadUUID, container *types.ViewThread, db *sqlx.DB) error {
	var err error

	err = db.Get(container, "SELECT category_uuid, name FROM app.threads WHERE thread_uuid = $1", thread_uuid)
	if err != nil {
		return errWrapper(err)
	}

	err = db.Get(container.Content.Posts, "SELECT user_uuid, post_uuid, date_added, date_last_edit, content  FROM app.posts WHERE thread_uuid = $1", thread_uuid)
	if err != nil {
		return errWrapper(err)
	}

	err = db.Get(container.Content.Users, "SELECT DISTINCT users.user_uuid, registration_date, nick_name, full_name, avatar256 FROM app.posts INNER JOIN app.users ON users.user_uuid = posts.user_uuid WHERE thread_uuid = $1", thread_uuid)
	if err != nil {
		return errWrapper(err)
	}

	return errWrapper(err)
}
func GetThreadsV1(category types.ModelCategoryUUID, container *types.ViewThreadsByCategory, db *sqlx.DB) error {
	var err = db.Get(container, "", category)
	return errWrapper(err)
}
