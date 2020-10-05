package inside

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/imakiri/playground/data/schema"
)

type Exec interface {
	ExecuteSQL() error
}

type Base struct {
	Data *schema.User
}

type GetUserV1 Base

func (d GetUserV1) ExecuteSQL() (err error) {
	q := goquDB.Select("name", "avatar").From(goqu.S("main").Table("users"))

	switch err := build("loginAndId", q, d.Data).(type) {
	case error:
		return err
	}

	b, e := q.ScanStruct(d.Data)
	if !b {
		return NotFoundError{}
	}

	return check(e)
}

type GetUserPassHashV1 Base

func (d GetUserPassHashV1) ExecuteSQL() (err error) {
	q := goquDB.Select("passHash").From(goqu.S("main").Table("users"))

	switch err := build("login", q, d.Data).(type) {
	case error:
		return err
	}

	b, e := q.ScanStruct(d.Data)
	if !b {
		return NotFoundError{}
	}

	return check(e)
}

type CreateUserV1 Base

func (d CreateUserV1) ExecuteSQL() (err error) {
	q := goquDB.Insert("users").Rows(d.Data)

	switch err := build("loginAndPassHash", q, d.Data).(type) {
	case error:
		return err
	}

	_, err = q.Executor().Exec()
	return check(err)
}

type DeleteUserV1 Base

func (d DeleteUserV1) ExecuteSQL() (err error) {
	q := goquDB.Delete("users")

	switch err := build("loginAndId", q, d.Data).(type) {
	case error:
		return err
	}

	re, e := q.Executor().Exec()
	if e != nil {
		return InternalServiceError{BaseError(e.Error())}
	}

	if te, _ := re.RowsAffected(); te == 0 {
		return NoUserToDelete{}
	}
	return check(e)
}
