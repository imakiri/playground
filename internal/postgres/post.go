package postgres

type ViewPostCreate struct {
	UserUUID   string
	ThreadUUID string
	Content    string
}

func PostCreate(conn Connection, container ViewPostCreate) error {
	return nil
}
func PostUpdate(conn Connection, postUUID string, container ViewPostCreate) error {
	return nil
}
func PostDelete(conn Connection, postUUID string) error {
	return nil
}
