package postgres

import (
	"github.com/imakiri/gorum/internal/data"
)

func UserGetProfile(conn Connection, uuid string, container *data.ViewUserProfile) error {
	return conn.db.Get(container, "SELECT registration_date, nick_name, full_name, avatar512 FROM main.app.users WHERE user_uuid = $1", uuid)
}
func UserUpdateProfile(conn Connection, uuid string, container data.ViewUserProfileUpdate) error {
	var tx, err = conn.db.Begin()
	if err != nil {
		return err
	}

	if container.Avatar != nil {
		_, err = tx.Exec("UPDATE app.users SET avatar512 = $2, avatar256 = $3, avatar128 = $4 WHERE user_uuid = $1", uuid, container.Avatar.Avatar512, container.Avatar.Avatar256, container.Avatar.Avatar128)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return e
			}
			return err
		}
	}
	if container.FullName != nil {
		_, err = tx.Exec("UPDATE app.users SET full_name = $2 WHERE user_uuid = $1", uuid, container.FullName)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return e
			}
			return err
		}
	}
	if container.NickName != nil {
		_, err = tx.Exec("UPDATE app.users SET nick_name = $2 WHERE user_uuid = $1", uuid, container.NickName)
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return e
			}
			return err
		}
	}

	err = tx.Commit()
	return err
}
