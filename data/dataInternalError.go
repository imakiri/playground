package data

import (
	"github.com/go-sql-driver/mysql"
)

type ERROR string

func (b ERROR) Error() string {
	return string(b)
}

type InternalServiceError struct{ ERROR }
type IncorrectArgumentError struct{ ERROR }
type NotFoundError struct{ ERROR }
type AlreadyExistError struct{ ERROR }

func check(err error) error {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return AlreadyExistError{ERROR(e.Error())}
		case 1048:
			return IncorrectArgumentError{ERROR(e.Error())}
		default:
			return InternalServiceError{ERROR(e.Error())}
		}
	case error:
		break
	default:
		return err
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return NotFoundError{ERROR(st)}
	default:
		return InternalServiceError{ERROR(st)}
	}
}

func checkRequest(id uint, login string) (err error) {
	switch {
	case id == 0 && login == "":
		return IncorrectArgumentError{ERROR("null id and login")}
	case id != 0 && login != "":
		return IncorrectArgumentError{ERROR("can accept only id or login, not both")}
	default:
		return nil
	}
}
