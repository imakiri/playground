package inside

import (
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/data/schema"
)

type checkerImp struct{}

func (checkerImp) createUser(u *schema.User, f userFunc) error {
	if u.Name == "" || u.Login == "" || u.Avatar == nil || u.PassHash == nil {
		return IncorrectArgumentError{"Invalid/null argument"}
	}

	if err := f(u); err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok {
			switch me.Number {
			case 1062:
				return IncorrectArgumentError{"User already exist"}
			default:
				return InternalServiceError{err.Error()}
			}
		} else {
			return InternalServiceError{err.Error()}
		}
	} else {
		return nil
	}
}

func (checkerImp) getUser(u *schema.User, fId userFunc, fLogin userFunc) error {
	if u.Login != "" && u.Id != 0 {
		return IncorrectArgumentError{"Couldn't accept both login and id as a parameter"}
	}

	if u.Login == "" && u.Id == 0 {
		return IncorrectArgumentError{"Invalid/null argument"}
	}

	if u.Id != 0 {
		if err := fId(u); err != nil {
			switch e := err.Error(); e {
			case "sql: no rows in result set":
				return NotFoundError{e}
			default:
				return InternalServiceError{e}
			}
		} else {
			return nil
		}
	}

	if u.Login != "" {
		if err := fLogin(u); err != nil {
			switch e := err.Error(); e {
			case "sql: no rows in result set":
				return NotFoundError{e}
			default:
				return InternalServiceError{e}
			}
		} else {
			return nil
		}
	}

	return InternalServiceError{}
}

type IncorrectArgumentErrorInt error

type InternalServiceError struct {
	err string
}

func (e InternalServiceError) Error() string {
	return e.err
}

type IncorrectArgumentError struct {
	err string
}

func (e IncorrectArgumentError) Error() string {
	return e.err
}

type NotFoundError struct {
	err string
}

func (e NotFoundError) Error() string {
	return e.err
}
