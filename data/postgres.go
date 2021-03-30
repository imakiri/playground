package data

import (
	"database/sql"
	"github.com/aidarkhanov/nanoid"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/types"
	"runtime"
	"strings"
	"time"
)

func funcName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}

// Wrapper for raw sql/sqlx/pgx error strings. Will panic if err == nil
func errrapper(err error) erres.Error {
	var f = funcName()

	switch {
	case err == nil:
		panic(f + ": nil error")
	case err == sql.ErrTxDone:
		return erres.InternalServiceError.Extend()
	}

	var e = err.Error()

	switch {
	case strings.Contains(e, "sqlx.bindNamedMapper: unsupported map type:"):
		return erres.InternalServiceError.Extend()
	default:
		return erres.JustError.Extend().AddRoute(f).AddDescription(e)
	}
}

func postgres_AddCookie_v1(uuid types.ModelUserUUID, cookie types.ViewCookieByUUID, p ServicePostgres) error {
	var c types.ModelCookie

	if uuid == "" {
		c.UUID = types.ModelUserUUID(nanoid.New())
	}

	c.Key = cookie.Key
	c.PemID = cookie.PemID
	c.ExpirationDate = cookie.ExpirationDate

	var _, err = p.db.NamedQuery("INSERT INTO main.auth.cookie VALUES (:key, :uuid, :pemid, :expirationDate)", c)
	return err
}
func postgres_GetCookie_v1(key types.ModelCookieKey, container *types.ViewCookieByUUID, p ServicePostgres) error {
	return p.db.Get(container, "SELECT uuid, pemid, expiration_date FROM main.auth.cookie WHERE key = $1", key)
}
func postgres_DeleteCookie_v1(uuid types.ModelUserUUID, p ServicePostgres) error {
	var _, err = p.db.Exec("DELETE FROM main.auth.cookie WHERE uuid = $1", uuid)
	return err
}
func postgres_AddLogpass_v1(uuid types.ModelUserUUID, logpass types.ViewLogpassByUUID, p ServicePostgres) error {
	var l types.ModelLogpass

	if uuid == "" {
		l.UUID = types.ModelUserUUID(nanoid.New())
	}

	l.PemID = logpass.PemID
	l.Password = logpass.Password
	l.Login = logpass.Login

	var _, err = p.db.NamedQuery("INSERT INTO main.auth.logpass VALUES (:uuid, :login, :password, :pemid)", l)
	return err
}
func postgres_GetLogpass_v1(login types.ModelLogpassLogin, container *types.ViewLogpassByUUID, p ServicePostgres) error {
	return p.db.Get(container, "SELECT uuid, pemid, password FROM main.auth.logpass WHERE login = $1", login)
}
func postgres_DeleteLogpass_v1(uuid types.ModelUserUUID, p ServicePostgres) error {
	var _, err = p.db.Exec("DELETE FROM main.auth.logpass WHERE uuid = $1", uuid)
	return err
}

func postgres_GetUserProfile_v1(uuid types.ModelUserUUID, container *types.ViewUserProfile, p ServicePostgres) error {
	var err = p.db.Get(container, "SELECT registration_date, nick_name, full_name, avatar512 FROM main.app.users", uuid)
	return errrapper(err)
}
func postgres_UpdateUserProfile_v1(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate, p ServicePostgres) error {
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
func postgres_CreateThread_v1(container types.ViewThreadCreate, p ServicePostgres) error {
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

	var _, err = p.db.NamedExec("INSERT INTO app.threads VALUES (:ThreadUUID, :CategoryUUID, :UserUUID, :Name, :DateAdded, :DateLastEdit, :Header)", thread)
	return errrapper(err)
}
func postgres_GetThread_v1(thread_uuid types.ModelThreadUUID, container *types.ViewThread, p ServicePostgres) error {
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
func postgres_GetThreads_v1(category types.ModelCategoryUUID, container *types.ViewThreadsByCategory, p ServicePostgres) error {
	var err = p.db.Get(container, "", category)
	return errrapper(err)
}
