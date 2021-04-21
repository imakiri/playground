package postgres

import (
	"github.com/imakiri/gorum/internal/types"
	"github.com/jmoiron/sqlx"
)

func CookieAdd(cookie types.ModelCookie, db *sqlx.DB) error {
	var _, err = db.NamedQuery("INSERT INTO main.auth.cookie VALUES (:Key, :UserUUID, :PemID, :ExpirationDate)", cookie)
	return errWrapper(err)
}
func CookieGet(key types.ModelCookieKey, container *types.ViewCookie, db *sqlx.DB) error {
	var err = db.Get(container, "SELECT user_uuid, pemid, expiration_date FROM main.auth.cookie WHERE key = $1", key)
	return errWrapper(err)
}
func CookieDelete(userUUID types.ModelUserUUID, db *sqlx.DB) error {
	var _, err = db.Exec("DELETE FROM main.auth.cookie WHERE user_uuid = $1", userUUID)
	return errWrapper(err)
}
