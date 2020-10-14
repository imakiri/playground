package core

type Execute interface {
	SQL()
}

type RequestData interface {
	Data()
}

// DataInternalMain Data Entities
type DataInternalMainGetUser_1 struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
	} `json:"request"`
	Response struct {
		DataInternalMainUserAvatar
		DataInternalMainUserName
	} `json:"response"`
	SQLfunc func(e *DataInternalMainGetUser_1)
}
type DataInternalMainGetUserPassHash_1 struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserLogin
	} `json:"request"`
	Response struct {
		DataInternalMainUserPassHash
	} `json:"response"`
	SQLfunc func(e *DataInternalMainGetUserPassHash_1)
}
type DataInternalMainCreateUser_1 struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserLogin
		DataInternalMainUserAvatar
		DataInternalMainUserName
		DataInternalMainUserPassHash
	} `json:"request"`
	SQLfunc func(e *DataInternalMainCreateUser_1)
}
type DataInternalMainDeleteUser_1 struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
	} `json:"request"`
	SQLfunc func(e *DataInternalMainDeleteUser_1)
}
type DataInternalMainUpdateUser_1 struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
		DataInternalMainUserName
		DataInternalMainUserAvatar
	} `json:"request"`
	SQLfunc func(e *DataInternalMainUpdateUser_1)
}
type DataInternalMainUpdateUserPassHash_1 struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserLogin
		DataInternalMainUserPassHash
	} `json:"request"`
	SQLfunc func(e *DataInternalMainUpdateUserPassHash_1)
}

// DataInternalMain General SQL Method Impl
func (e *DataInternalMainGetUser_1) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainGetUserPassHash_1) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainCreateUser_1) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainDeleteUser_1) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainUpdateUser_1) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainUpdateUserPassHash_1) SQL() {
	e.SQLfunc(e)
}
