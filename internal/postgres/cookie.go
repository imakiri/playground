package postgres

type ModelCookie struct {
	UserUUID string
	ViewCookie
}

type ViewCookie struct {
	PemID          int16
	Key            string
	ExpirationDate int64
}

func CookieAdd(conn Connection, cookie ModelCookie) error {
	var _, err = conn.db.NamedQuery("INSERT INTO main.auth.cookie VALUES (:Key, :UserUUID, :PemID, :ExpirationDate)", cookie)
	return err
}
func CookieGet(conn Connection, key string, container *ViewCookie) error {
	return conn.db.Get(container, "SELECT user_uuid, pemid, expiration_date FROM main.auth.cookie WHERE key = $1", key)
}
func CookieDelete(conn Connection, userUUID string) error {
	var _, err = conn.db.Exec("DELETE FROM main.auth.cookie WHERE user_uuid = $1", userUUID)
	return err
}
