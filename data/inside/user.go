package inside

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/imakiri/playground/data/schema"
)

type Exec interface {
	ExecuteSQL() error
}

type BaseUser struct {
	Data *schema.User
}

type GetUserV1 BaseUser

func (d GetUserV1) ExecuteSQL() (err error) {
	q := goquDB.Select("name", "avatar").From(goqu.S("main").Table("users"))

	switch e := build("loginAndId", q, d.Data).(type) {
	case error:
		return e
	}

	b, e := q.ScanStruct(d.Data)
	if !b {
		return NotFoundError{}
	}

	return check(e)
}

type GetUserPassHashV1 BaseUser

func (d GetUserPassHashV1) ExecuteSQL() (err error) {
	q := goquDB.Select("passHash").From(goqu.S("main").Table("users"))

	switch e := build("login", q, d.Data).(type) {
	case error:
		return e
	}

	b, e := q.ScanStruct(d.Data)
	if !b {
		return NotFoundError{}
	}

	return check(e)
}

type CreateUserV1 BaseUser

func (d CreateUserV1) ExecuteSQL() (err error) {
	q := goquDB.Insert("users").Rows(d.Data)

	switch e := build("loginAndPassHash", q, d.Data).(type) {
	case error:
		return e
	}

	_, err = q.Executor().Exec()
	return check(err)
}

type DeleteUserV1 BaseUser

func (d DeleteUserV1) ExecuteSQL() (err error) {
	q := goquDB.Delete("users")

	switch e := build("loginAndId", q, d.Data).(type) {
	case error:
		return e
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

//type UpdateUserV1 BaseUser
////
////func (d UpdateUserV1) ExecuteSQL() (err error) {
////	q := goquDB.Update("users")
////
////	switch e := build("update", q, d.Data).(type) {
////	case error:
////		return e
////	}
////	////////////////////////////////////////////
////	return nil
////}

type GetUserV2 BaseUser

func (d GetUserV2) ExecuteSQL() (err error) {
	t := new([]schema.User)
	err = sqlxDB.Select(t, "SELECT name FROM main.users WHERE login = ? OR  id = ?", d.Data.Login, d.Data.Id)

	switch {
	case len(*t) > 1:
		return IncorrectArgumentError{}
	default:
		*d.Data = (*t)[0]
		return check(err)
	}
}
