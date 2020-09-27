package internal

import "github.com/imakiri/playground/data/schema"

func (consecutive) UpdateUserAvatar(u schema.User) (err error) {
	_, err = main.Exec("UPDATE main.users SET avatar = ? WHERE login = ?", u.Avatar, u.Login)
	return
}

func (consecutive) UpdateUserName(u schema.User) (err error) {
	_, err = main.Exec("UPDATE main.users SET name = ? WHERE login = ?", u.Name, u.Login)
	return
}
