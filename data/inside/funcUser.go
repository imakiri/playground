package inside

import (
	"github.com/imakiri/playground/data/schema"
)

var c userCheckerImp

type userChecker interface {
	userLI(u *schema.User, fId userFunc, fLogin userFunc) error
	userL(u *schema.User, fLogin userFunc) error
	createUser(u *schema.User, f userFunc) error
	updateUser(u *schema.User, fName userFunc, fAvatar userFunc, fLogin userFunc) error
}

type userFunc func(u *schema.User) (err error)

func (R) GetUser(u *schema.User) (err error) {
	fId := func(u *schema.User) (err error) {
		err = main.Get(u, "SELECT name, avatar FROM main.users WHERE id = ?", u.Id)
		return
	}

	fLogin := func(u *schema.User) (err error) {
		err = main.Get(u, "SELECT name, avatar FROM main.users WHERE login = ?", u.Login)
		return
	}

	return userChecker.userLI(c, u, fId, fLogin)
}

func (R) GetUserPassHash(u *schema.User) (err error) {
	f := func(u *schema.User) (err error) {
		err = main.Get(u, "SELECT passHash FROM main.users WHERE login = ?", u.Login)
		return
	}

	return userChecker.userL(c, u, f)
}

func (R) CreateUser(u *schema.User) (err error) {
	f := func(u *schema.User) (err error) {
		_, err = main.Exec("INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)",
			u.Login, u.Name, u.Avatar, u.PassHash)
		return
	}

	return userChecker.createUser(c, u, f)
}

func (R) DeleteUser(u *schema.User) (err error) {
	fId := func(u *schema.User) (err error) {
		_, err = main.Exec("DELETE FROM main.users WHERE id = ?", u.Id)
		return
	}

	fLogin := func(u *schema.User) (err error) {
		_, err = main.Exec("DELETE FROM main.users WHERE login = ?", u.Login)
		return
	}

	return userChecker.userLI(c, u, fId, fLogin)
}

func (R) UpdateUser(u *schema.User) (err error) {
	//fName := func(u *schema.User, fName userFunc, fAvatar userFunc, fLogin userFunc) (err error) {
	//	_, err = main.Exec("UPDATE main.users SET name = ? WHERE login = ?", u.Name, u.Login)
	//	return
	//}
	//
	//fAvatar := func() {
	//	_, err = main.Exec("UPDATE main.users SET avatar = ? WHERE login = ?", u.Avatar, u.Login)
	//	return
	//}
	//
	//

	return
}

func (R) UpdateUserName(u *schema.User) (err error) {
	_, err = main.Exec("UPDATE main.users SET name = ? WHERE login = ?", u.Name, u.Login)
	return
}
