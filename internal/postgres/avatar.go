package postgres

type ViewUserAvatar struct {
	Avatar512 []byte
	Avatar256 []byte
	Avatar128 []byte
}

func AvatarGet128(conn Connection, userUUID string, dest []byte) error {
	return conn.db.Get(dest, "SELECT avatar128 FROM main.content.avatars WHERE user_uuid = $1", userUUID)
}
func AvatarGet256(conn Connection, userUUID string, dest []byte) error {
	return conn.db.Get(dest, "SELECT avatar256 FROM main.content.avatars WHERE user_uuid = $1", userUUID)
}
func AvatarGet512(conn Connection, userUUID string, dest []byte) error {
	return conn.db.Get(dest, "SELECT avatar512 FROM main.content.avatars WHERE user_uuid = $1", userUUID)
}
func AvatarSet(conn Connection, update bool, userUUID string, avatar ViewUserAvatar) error {
	var err error
	var container = struct {
		UserUUID string
		ViewUserAvatar
	}{
		UserUUID:       userUUID,
		ViewUserAvatar: avatar,
	}
	if update {
		_, err = conn.db.NamedExec("UPDATE main.content.avatars SET avatar512 = :Avatar512, avatar256 = :Avatar256, avatar128 = :Avatar128 WHERE user_uuid = :UserUUID", container)
	} else {
		_, err = conn.db.NamedExec("INSERT INTO main.content.avatars VALUES (:UserUUID, :Avatar512, :Avatar256, :Avatar128)", container)
	}
	return err
}
