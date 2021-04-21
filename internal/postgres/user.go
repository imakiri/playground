package postgres

import (
	"github.com/imakiri/gorum/internal/types"
	"github.com/jmoiron/sqlx"
)

func UserGetProfile(uuid types.ModelUserUUID, container *types.ViewUserProfile, db *sqlx.DB) error {
	var err = db.Get(container, "SELECT registration_date, nick_name, full_name, avatar512 FROM main.app.users WHERE user_uuid = $1", uuid)
	return errWrapper(err)
}
func UserUpdateProfile(uuid types.ModelUserUUID, container types.ViewUserProfileUpdate, db *sqlx.DB) error {
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
