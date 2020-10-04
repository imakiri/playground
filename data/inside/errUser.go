package inside

import (
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/data/schema"
)

func check(err error) error {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return UserAlreadyExistError{e.Error()}
		default:
			return InternalServiceError{e.Error()}
		}
	case error:
		break
	default:
		return err
	}

	switch e := err.Error(); e {
	case "sql: no rows in result set":
		return NotFoundError{e}
	default:
		return InternalServiceError{e}
	}
}

type userCheckerImp struct{}

func (userCheckerImp) createUser(u *schema.User, f userFunc) error {
	if u.Name == "" || u.Login == "" || u.Avatar == nil || u.PassHash == nil {
		return IncorrectArgumentError{"Invalid/null argument"}
	}
	return check(f(u))
}

func (userCheckerImp) userLI(u *schema.User, fId userFunc, fLogin userFunc) (err error) {
	switch {
	case u.Login != "" && u.Id != 0:
		return IncorrectArgumentError{"Couldn't accept both login and id as a parameter"}
	case u.Login == "" && u.Id == 0:
		return IncorrectArgumentError{"Invalid/null argument"}
	case u.Id != 0:
		err = fId(u)
	case u.Login != "":
		err = fLogin(u)
	}

	return check(err)
}

func (userCheckerImp) userL(u *schema.User, f userFunc) (err error) {
	if u.Login == "" {
		return IncorrectArgumentError{"Invalid/null argument"}
	}
	return check(f(u))
}

func (userCheckerImp) updateUser(u *schema.User, fName userFunc, fAvatar userFunc, fLogin userFunc) (err error) {
	switch {
	case u.Login != "" && u.Id != 0:
		return IncorrectArgumentError{"Couldn't accept both login and id as a parameter"}
	case u.Login == "" && u.Id == 0:
		return IncorrectArgumentError{"Invalid/null argument"}
	case u.Name != "" && fName != nil:

	}

	if u.Name != "" && fName != nil {
		if err := fName; err != nil {

		}
	}
	return
}
