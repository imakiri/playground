package inside

import (
	"github.com/imakiri/playground/data/schema"
)

type checkerImp struct{}

func (checkerImp) createUser(u *schema.User) error {
	c := u.Name == "" || u.Login == "" || u.Avatar == nil || u.PassHash == nil
	if c {
		return IncorrectArgumentError{}
	}
	return nil
}

func (checkerImp) user(u *schema.User, fId userFunc, fLogin userFunc) error {
	var c bool

	c = u.Login != "" && u.Id != 0
	if c {
		return IncorrectArgumentError{}
	}

	c = u.Login == "" && u.Id == 0
	if c {
		return IncorrectArgumentError{}
	}

	if u.Id != 0 {
		err := fId(u)
		if err != nil && err.Error() == "sql: no rows in result set" {
			return NotFoundError{}
		}

		return err
	}

	if u.Login != "" {
		err := fLogin(u)
		if err != nil && err.Error() == "sql: no rows in result set" {
			return NotFoundError{}
		}

		return err
	}

	return NotFoundError{}
}

type IncorrectArgumentError struct {
	err string
}

func (IncorrectArgumentError) Error() string {
	return ""
}

type NotFoundError struct {
	err string
}

func (NotFoundError) Error() string {
	return ""
}
