package postgres

import (
	"github.com/imakiri/gorum/internal/types"
	"github.com/jmoiron/sqlx"
)

func LogpassAdd(logpass types.ModelLogpass, db *sqlx.DB) error {
	var _, err = db.NamedQuery("INSERT INTO main.auth.logpass VALUES (:UserUUID, :Login, :Password, :PemID)", logpass)
	return errWrapper(err)
}
func LogpassGetWithLogin(login types.ModelLogpassLogin, container *types.ViewLogpass, db *sqlx.DB) error {
	var err = db.Get(container, "SELECT user_uuid, pemid, password FROM main.auth.logpass WHERE login = $1", login)
	return errWrapper(err)
}
func LogpassDelete(userUUID types.ModelUserUUID, db *sqlx.DB) error {
	var _, err = db.Exec("DELETE FROM main.auth.logpass WHERE user_uuid = $1", userUUID)
	return errWrapper(err)
}
