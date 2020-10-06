package inside

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/data/schema"
)

func check(err error) error {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return UserAlreadyExistError{BaseError(e.Error())}
		default:
			return InternalServiceError{BaseError(e.Error())}
		}
	case error:
		break
	default:
		return err
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return NotFoundError{BaseError(st)}
	default:
		return InternalServiceError{BaseError(st)}
	}
}

func build(c string, sun interface{}, m interface{}) (err error) {
	switch d := (m).(type) {
	case *schema.User:
		switch c {
		case "loginAndId":
			switch {
			case d.Id != 0 && d.Login != "":
				return IncorrectArgumentError{"build: uncertain argument"}
			case d.Id == 0 && d.Login == "":
				return IncorrectArgumentError{"build: null argument"}
			case d.Id != 0:
				switch s := sun.(type) {
				case *goqu.SelectDataset:
					*s = *s.Where(goqu.C("id").Eq(d.Id))
					return nil
				case *goqu.DeleteDataset:
					*s = *s.Where(goqu.C("id").Eq(d.Id))
					return nil
				default:
					return InternalServiceError{"build: invalid sun"}
				}
			case d.Login != "":
				switch s := sun.(type) {
				case *goqu.SelectDataset:
					*s = *s.Where(goqu.C("login").Eq(d.Login))
					return nil
				case *goqu.DeleteDataset:
					*s = *s.Where(goqu.C("id").Eq(d.Login))
					return nil
				default:
					return InternalServiceError{"build: invalid sun"}
				}
			default:
				return InternalServiceError{"build: nil argument pointer"}
			}
		case "login":
			switch {
			case d.Login == "":
				return IncorrectArgumentError{"build: null argument"}
			case d.Login != "":
				switch s := sun.(type) {
				case *goqu.SelectDataset:
					*s = *s.Where(goqu.C("login").Eq(d.Login))
					return nil
				default:
					return InternalServiceError{"build: invalid sun"}
				}
			default:
				return InternalServiceError{"build: nil argument pointer"}
			}
		case "loginAndPassHash":
			switch {
			case d.Login == "" && d.PassHash == nil:
				return IncorrectArgumentError{"build: null argument"}
			case d.Login != "" && d.PassHash != nil:
				return nil
			default:
				return InternalServiceError{"build: nil argument pointer"}
			}
		case "update":
			switch {
			case d.Id != 0 && d.Login != "":
				return IncorrectArgumentError{"build: uncertain argument"}
			case d.Id == 0 && d.Login == "":
				return IncorrectArgumentError{"build: null argument"}
			case d.Id != 0:
				switch s := sun.(type) {
				case *goqu.SelectDataset:
					*s = *s.Where(goqu.C("id").Eq(d.Id))
					return nil
				case *goqu.DeleteDataset:
					*s = *s.Where(goqu.C("id").Eq(d.Id))
					return nil
				default:
					return InternalServiceError{"build: invalid sun"}
				}
			case d.Login != "":
				switch s := sun.(type) {
				case *goqu.SelectDataset:
					*s = *s.Where(goqu.C("login").Eq(d.Login))
					return nil
				case *goqu.DeleteDataset:
					*s = *s.Where(goqu.C("id").Eq(d.Login))
					return nil
				default:
					return InternalServiceError{"build: invalid sun"}
				}
			default:
				return InternalServiceError{"build: nil argument pointer"}
			}
		default:
			return InternalServiceError{"build: invalid build type"}
		}
	default:
		return InternalServiceError{"build: invalid data type"}
	}
}
