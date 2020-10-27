package data

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/core"
	"github.com/jmoiron/sqlx"
)

var Salt string

func init() {}

type RequestInternalMainDeleteUser struct{}
type RequestInternalMainGetUser struct{}
type RequestInternalMainGetUserPassHash struct{}
type RequestInternalMainCreateUser struct{}
type RequestInternalMainUpdateUser struct{}
type RequestInternalMainUpdateUserPassHash struct{}

func (RequestInternalMainDeleteUser) Data()         {}
func (RequestInternalMainGetUser) Data()            {}
func (RequestInternalMainGetUserPassHash) Data()    {}
func (RequestInternalMainCreateUser) Data()         {}
func (RequestInternalMainUpdateUser) Data()         {}
func (RequestInternalMainUpdateUserPassHash) Data() {}

func NewRequest(request core.RequestData) core.Execute {
	switch request.(type) {
	case RequestInternalMainCreateUser:
		return &core.DataInternalMainCreateUser{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserLogin
				core.DataInternalMainUserAvatar
				core.DataInternalMainUserName
				core.DataInternalMainUserPassHash
			}{},
			SQLfunc: dataInternalMainCreateUser,
		}
	case RequestInternalMainGetUser:
		return &core.DataInternalMainGetUser{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserId
				core.DataInternalMainUserLogin
			}{},
			Response: struct {
				core.DataInternalMainUserAvatar
				core.DataInternalMainUserName
			}{},
			SQLfunc: dataInternalMainGetUser,
		}
	case RequestInternalMainGetUserPassHash:
		return &core.DataInternalMainGetUserPassHash{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserLogin
			}{},
			Response: struct {
				core.DataInternalMainUserPassHash
			}{},
			SQLfunc: dataInternalMainGetPassHash,
		}
	case RequestInternalMainUpdateUser:
		return &core.DataInternalMainUpdateUser{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserId
				core.DataInternalMainUserLogin
				core.DataInternalMainUserName
				core.DataInternalMainUserAvatar
			}{},
			SQLfunc: dataInternalMainUpdateUser,
		}
	case RequestInternalMainUpdateUserPassHash:
		return &core.DataInternalMainUpdateUserPassHash{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserLogin
				core.DataInternalMainUserPassHash
			}{},
			SQLfunc: dataInternalMainUpdateUserPassHash,
		}
	case RequestInternalMainDeleteUser:
		return &core.DataInternalMainDeleteUser{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserId
				core.DataInternalMainUserLogin
			}{},
			SQLfunc: dataInternalMainDeleteUser,
		}
	default:
		return nil
	}
}

func prepareQueryErrorCheck(p core.StatusPointer, funcName string, err error) bool {
	s := p.StatusPtr()
	if err != nil {
		s.Type_ = core.StatusType2StrCaster(core.SerializationError{})
		s.Value = fmt.Sprintf("%s: %s", funcName, err.Error())
		return false
	} else {
		s.Type_ = core.StatusType2StrCaster(core.StatusOk{})
		return true
	}
}

func executeQueryErrorCheck(e core.StatusPointer, err error) bool {
	s := e.StatusPtr()
	if err != nil {
		s.Value = err.Error()
		s.Type_ = errTypeCast(err)
		return false
	} else {
		s.Type_ = core.StatusType2StrCaster(core.StatusOk{})
		return true
	}

}

func errTypeCast(err error) (type_ string) {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return core.StatusType2StrCaster(core.DataAlreadyExistError{})
		case 1048:
			return core.StatusType2StrCaster(core.DataIncorrectArgumentError{})
		default:
			return core.StatusType2StrCaster(core.DataInternalServiceError{})
		}
	case error:
		break
	default:
		return core.StatusType2StrCaster(core.UnknownError{})
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return core.StatusType2StrCaster(core.DataNotFoundError{})
	default:
		return core.StatusType2StrCaster(core.DataInternalServiceError{})
	}
}

const query_dataInternalMainCreateUser = "INSERT INTO main.users (login, name, avatar, pass_hash) VALUES (:login, :name, :avatar, :pass_hash)"
const query_dataInternalMainGetUser = "SELECT name, avatar FROM main.users WHERE login = :login"
const query_dataInternalMainGetPassHash = "SELECT pass_hash FROM main.users WHERE login = :login"
const query_dataInternalMainUpdateUser = "UPDATE main.users SET name = :name, avatar = :avatar WHERE login = :login"
const query_dataInternalMainUpdateUserPassHash = "UPDATE main.users SET pass_hash = :pass_hash WHERE login = :login"
const query_dataInternalMainDeleteUser = "DELETE FROM main.users WHERE login = :login"

func dataInternalMainCreateUser(e *core.DataInternalMainCreateUser) {
	q, args, err := sqlx.Named(query_dataInternalMainCreateUser, e.Request)
	if !prepareQueryErrorCheck(e, "dataInternalMainCreateUser", err) {
		return
	}

	_, err = e.SQLX_DB.Exec(q, args...)
	if !executeQueryErrorCheck(e, err) {
		return
	}
}
func dataInternalMainGetUser(e *core.DataInternalMainGetUser) {
	q, args, err := sqlx.Named(query_dataInternalMainGetUser, e.Request)

	if !prepareQueryErrorCheck(e, "dataInternalMainGetUser", err) {
		return
	}

	q = e.SQLX_DB.Rebind(q)
	err = e.SQLX_DB.Get(&e.Response, q, args...)
	if !executeQueryErrorCheck(e, err) {
		return
	}
}
func dataInternalMainGetPassHash(e *core.DataInternalMainGetUserPassHash) {
	q, args, err := sqlx.Named(query_dataInternalMainGetPassHash, e.Request)

	if !prepareQueryErrorCheck(e, "dataInternalMainGetPassHash", err) {
		return
	}

	q = e.SQLX_DB.Rebind(q)
	err = e.SQLX_DB.Get(&e.Response, q, args...)
	if !executeQueryErrorCheck(e, err) {
		return
	}
}
func dataInternalMainUpdateUser(e *core.DataInternalMainUpdateUser) {
	q, args, err := sqlx.Named(query_dataInternalMainUpdateUser, e.Request)
	if !prepareQueryErrorCheck(e, "dataInternalMainUpdateUser", err) {
		return
	}

	_, err = e.SQLX_DB.Exec(q, args...)
	if !executeQueryErrorCheck(e, err) {
		return
	}
}
func dataInternalMainUpdateUserPassHash(e *core.DataInternalMainUpdateUserPassHash) {
	q, args, err := sqlx.Named(query_dataInternalMainUpdateUserPassHash, e.Request)
	if !prepareQueryErrorCheck(e, "dataInternalMainUpdateUserPassHash", err) {
		return
	}

	_, err = e.SQLX_DB.Exec(q, args...)
	if !executeQueryErrorCheck(e, err) {
		return
	}
}
func dataInternalMainDeleteUser(e *core.DataInternalMainDeleteUser) {
	q, args, err := sqlx.Named(query_dataInternalMainDeleteUser, e.Request)
	if !prepareQueryErrorCheck(e, "dataInternalMainDeleteUser", err) {
		return
	}

	_, err = e.SQLX_DB.Exec(q, args...)
	if !executeQueryErrorCheck(e, err) {
		return
	}
}
