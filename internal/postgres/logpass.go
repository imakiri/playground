package postgres

type ModelLogpass struct {
	UserUUID string
	ViewLogpass
}

type ViewLogpass struct {
	PemID    int16
	Login    []byte
	Password []byte
}

func LogpassAdd(conn Connection, logpass ModelLogpass) error {
	var _, err = conn.db.NamedQuery("INSERT INTO main.auth.logpass VALUES (:UserUUID, :Login, :Password, :PemID)", logpass)
	return err
}
func LogpassGetWithLogin(conn Connection, login []byte, container *ViewLogpass) error {
	return conn.db.Get(container, "SELECT user_uuid, pemid, password FROM main.auth.logpass WHERE login = $1", login)
}
func LogpassDelete(conn Connection, userUUID string) error {
	var _, err = conn.db.Exec("DELETE FROM main.auth.logpass WHERE user_uuid = $1", userUUID)
	return err
}
