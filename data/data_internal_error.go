package data

import (
	"github.com/go-sql-driver/mysql"
)

type ERROR string

func (b ERROR) Error() string {
	return string(b)
}

type Internal_ERROR_InternalServiceError struct{ ERROR }
type Internal_ERROR_IncorrectArgument struct{ ERROR }
type Internal_ERROR_NotFound struct{ ERROR }
type Internal_ERROR_AlreadyExist struct{ ERROR }

func check(err error) error {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return Internal_ERROR_AlreadyExist{ERROR(e.Error())}
		case 1048:
			return Internal_ERROR_IncorrectArgument{ERROR(e.Error())}
		default:
			return Internal_ERROR_InternalServiceError{ERROR(e.Error())}
		}
	case error:
		break
	default:
		return err
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return Internal_ERROR_NotFound{ERROR(st)}
	default:
		return Internal_ERROR_InternalServiceError{ERROR(st)}
	}
}
