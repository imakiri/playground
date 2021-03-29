package data

import (
	"database/sql"
	"github.com/imakiri/erres"
	"runtime"
	"strings"
)

type (
	ModelUserUUID             string
	ModelUserPemID            int16
	ModelUserNickName         string
	ModelUserFullName         string
	ModelUserPosts            int32
	ModelUserAvatar512        []byte
	ModelUserAvatar256        []byte
	ModelUserAvatar128        []byte
	ModelUserRegistrationDate int64

	ModelDate    int64
	ModelContent string

	ModelPostUUID string

	ModelThreadUUID string
	ModelThreadName string

	ModelCategoryUUID string
	ModelCategoryName string

	ModelCookieKey            string
	ModelCookieExpirationDate int64

	ModelLogpassLogin    []byte
	ModelLogpassPassword []byte
)

type (
	ModelUser struct {
		UserUUID         ModelUserUUID
		RegistrationDate ModelUserRegistrationDate
		Nickname         ModelUserNickName
		Fullname         ModelUserFullName
		Avatar512        ModelUserAvatar512
		Avatar256        ModelUserAvatar256
		Avatar128        ModelUserAvatar128
	}
	ModelThread struct {
		ThreadUUID   ModelThreadUUID
		CategoryUUID ModelCategoryUUID
		UserUUID     ModelUserUUID
		Name         ModelThreadName
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Header       ModelContent
	}
	ModelPost struct {
		PostUUID     ModelPostUUID
		ThreadUUID   ModelThreadUUID
		UserUUID     ModelUserUUID
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Content      ModelContent
	}
	ModelCategory struct {
		CategoryUUID ModelCategoryUUID
		Name         ModelCategoryName
	}
	ModelCookie struct {
		Key            ModelCookieKey
		UUID           ModelUserUUID
		PemID          ModelUserPemID
		ExpirationDate ModelCookieExpirationDate
	}
	ModelLogpass struct {
		UUID     ModelUserUUID
		Login    ModelLogpassLogin
		Password ModelLogpassPassword
		PemID    ModelUserPemID
	}
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
