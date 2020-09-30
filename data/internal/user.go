package internal

import "github.com/imakiri/playground/data/schema"

func (r) GetUserById(id uint) (re schema.Re) {
	re.Err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE id = ?", id)
	return
}

func (r) GetUserByLogin(login string) (re schema.Re) {
	re.Err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE login = ?", login)
	return
}

func (r) GetUserPassHash(login string) (re schema.Re) {
	re.Err = main.Get(&re.User, "SELECT passHash FROM main.users WHERE login = ?", login)
	return
}

func (r) CreateUser(u schema.User) (err error) {
	_, err = main.Exec("INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)",
		u.Login, u.Name, u.Avatar, u.PassHash)
	return
}

func (r) DeleteUser(login string) (err error) {
	_, err = main.Exec("DELETE FROM main.users WHERE login = ?", login)
	return
}

func (r) UpdateUserAvatar(u schema.User) (err error) {
	_, err = main.Exec("UPDATE main.users SET avatar = ? WHERE login = ?", u.Avatar, u.Login)
	return
}

func (r) UpdateUserName(u schema.User) (err error) {
	_, err = main.Exec("UPDATE main.users SET name = ? WHERE login = ?", u.Name, u.Login)
	return
}
