package internal

import "github.com/imakiri/playground/data/schema"

func (consecutive) GetUserById(id uint) (re schema.Re) {
	re.Err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE id = ?", id)
	return
}

func (consecutive) GetUserByLogin(login string) (re schema.Re) {
	re.Err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE login = ?", login)
	return
}

func (consecutive) GetUserPassHash(login string) (re schema.Re) {
	re.Err = main.Get(&re.User, "SELECT passHash FROM main.users WHERE login = ?", login)
	return
}

func (concurrent) GetUserById(id uint, c chan schema.Re) {
	defer close(c)

	re := schema.Re{}
	re.Err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE id = ?", id)
	c <- re
}

func (concurrent) GetUserByLogin(login string, c chan schema.Re) {
	defer close(c)

	re := schema.Re{}
	re.Err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE login = ?", login)
	c <- re
}

func (concurrent) GetUserPassHash(login string, c chan schema.Re) {
	defer close(c)

	re := schema.Re{}
	re.Err = main.Get(&re.User, "SELECT passHash FROM main.users WHERE login = ?", login)
	c <- re
}
