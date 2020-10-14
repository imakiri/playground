package data

import (
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/core"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
)

var Salt string

func init() {
	f, err := ioutil.ReadFile("data/salt")
	if err != nil {
		panic(err)
	}

	Salt = string(f)
}

type RequestInternalMainDeleteUser_1 struct{}
type RequestInternalMainGetUser_1 struct{}
type RequestInternalMainGetUserPassHash_1 struct{}
type RequestInternalMainCreateUser_1 struct{}
type RequestInternalMainUpdateUser_1 struct{}
type RequestInternalMainUpdateUserPassHash_1 struct{}

func (RequestInternalMainDeleteUser_1) Data()         {}
func (RequestInternalMainGetUser_1) Data()            {}
func (RequestInternalMainGetUserPassHash_1) Data()    {}
func (RequestInternalMainCreateUser_1) Data()         {}
func (RequestInternalMainUpdateUser_1) Data()         {}
func (RequestInternalMainUpdateUserPassHash_1) Data() {}

func NewRequest(request core.RequestData) core.Execute {
	switch request.(type) {
	case RequestInternalMainDeleteUser_1:
		return &core.DataInternalMainDeleteUser_1{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserId
				core.DataInternalMainUserLogin
			}{},
			SQLfunc: requestInternalMainDeleteUser_1,
		}
	case RequestInternalMainCreateUser_1:
		return &core.DataInternalMainCreateUser_1{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserLogin
				core.DataInternalMainUserAvatar
				core.DataInternalMainUserName
				core.DataInternalMainUserPassHash
			}{},
			SQLfunc: requestInternalMainCreateUser_1,
		}
	case RequestInternalMainGetUserPassHash_1:
		return &core.DataInternalMainGetUserPassHash_1{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserLogin
			}{},
			Response: struct {
				core.DataInternalMainUserPassHash
			}{},
			SQLfunc: requestInternalMainGetUserPassHash_1,
		}
	case RequestInternalMainGetUser_1:
		return &core.DataInternalMainGetUser_1{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserId
				core.DataInternalMainUserLogin
			}{},
			Response: struct {
				core.DataInternalMainUserAvatar
				core.DataInternalMainUserName
			}{},
			SQLfunc: requestInternalMainGetUser_1,
		}
	case RequestInternalMainUpdateUser_1:
		return &core.DataInternalMainUpdateUser_1{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserId
				core.DataInternalMainUserLogin
				core.DataInternalMainUserName
				core.DataInternalMainUserAvatar
			}{},
			SQLfunc: requestInternalMainUpdateUser_1,
		}
	case RequestInternalMainUpdateUserPassHash_1:
		return &core.DataInternalMainUpdateUserPassHash_1{
			DataInternalMain: ConnectionInternalMain,
			Request: struct {
				core.DataInternalMainUserLogin
				core.DataInternalMainUserPassHash
			}{},
			SQLfunc: requestInternalMainUpdateUserPassHash_1,
		}
	default:
		return nil
	}
}

func check(err error) core.Err {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return core.NewError(core.DataAlreadyExistError{}, e.Error())
		case 1048:
			return core.NewError(core.DataIncorrectArgumentError{}, e.Error())
		default:
			return core.NewError(core.DataInternalServiceError{}, e.Error())
		}
	case error:
		break
	default:
		return nil
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return core.NewError(core.DataNotFoundError{}, st)
	default:
		return core.NewError(core.DataInternalServiceError{}, st)
	}
}

func checkRequest(id uint, login string) core.Err {
	switch {
	case id == 0 && login == "":
		return core.NewError(core.DataIncorrectArgumentError{}, "null id and login")
	case id != 0 && login != "":
		return core.NewError(core.DataIncorrectArgumentError{}, "can accept only id or login, not both")
	default:
		return nil
	}
}
