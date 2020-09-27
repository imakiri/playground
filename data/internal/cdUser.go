package internal

import "github.com/imakiri/playground/data/schema"

func (consecutive) CreateUser(u schema.User) (err error) {
	_, err = main.Exec("INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)",
		u.Login, u.Name, u.Avatar, u.PassHash)
	return
}

func (consecutive) DeleteUser(login string) (err error) {
	_, err = main.Exec("DELETE FROM main.users WHERE login = ?", login)
	return
}
