package postgres

import (
	"github.com/imakiri/gorum/types"
	"github.com/jmoiron/sqlx"
)

func AvatarGet128(db *sqlx.DB, userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) error {
	var err = db.Get(container, "SELECT avatar128 FROM main.content.avatars WHERE user_uuid = $1", userUUID)
	return errWrapper(err)
}
func AvatarGet256(db *sqlx.DB, userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) error {
	var err = db.Get(container, "SELECT avatar256 FROM main.content.avatars WHERE user_uuid = $1", userUUID)
	return errWrapper(err)
}
func AvatarGet512(db *sqlx.DB, userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) error {
	var err = db.Get(container, "SELECT avatar512 FROM main.content.avatars WHERE user_uuid = $1", userUUID)
	return errWrapper(err)
}
func AvatarSet(db *sqlx.DB, update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	var err error
	var container = struct {
		UserUUID types.ModelUserUUID
		types.ViewUserAvatar
	}{
		UserUUID:       userUUID,
		ViewUserAvatar: avatar,
	}
	if update {
		_, err = db.NamedExec("UPDATE main.content.avatars SET avatar512 = :Avatar512, avatar256 = :Avatar256, avatar128 = :Avatar128 WHERE user_uuid = :UserUUID", container)
	} else {
		_, err = db.NamedExec("INSERT INTO main.content.avatars VALUES (:UserUUID, :Avatar512, :Avatar256, :Avatar128)", container)
	}
	return errWrapper(err)
}
