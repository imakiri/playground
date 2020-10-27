package core

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Execute interface {
	SQL()
}

type RequestData interface {
	Data()
}

// DataInternalMain Data Entity
type DataInternalMain struct {
	SQLX_DB *sqlx.DB
}

// DataInternalMainUser Fields
type DataInternalMainUserId struct {
	Id uint `db:"id" json:"dataInternalMainUserId"`
}
type DataInternalMainUserLogin struct {
	Login string `db:"login" json:"dataInternalMainUserLogin"`
}
type DataInternalMainUserAvatar struct {
	Avatar []byte `db:"avatar" json:"dataInternalMainUserAvatar"`
}
type DataInternalMainUserName struct {
	Name string `db:"name" json:"dataInternalMainUserName"`
}
type DataInternalMainUserPassHash struct {
	PassHash []byte `db:"pass_hash" json:"dataInternalMainUserPassHash"`
}

// DataInternalMainData Fields
type DataInternalMainDataId struct {
	Id uint `db:"id" json:"dataInternalMainDataId"`
}
type DataInternalMainDataDate struct {
	Date time.Time `db:"date" json:"dataInternalMainDataDate"`
}
type DataInternalMainDataQuery struct {
	Query string `db:"query" json:"dataInternalMainDataQuery"`
}
type DataInternalMainDataPic struct {
	Pic []byte `db:"pic" json:"dataInternalMainDataPic"`
}
type DataInternalMainDataUserId struct {
	UserId uint `db:"user_id" json:"dataInternalMainDataUserId"`
}

// DataInternalMain Data Entities
type DataInternalMainCreateUser struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserLogin
		DataInternalMainUserAvatar
		DataInternalMainUserName
		DataInternalMainUserPassHash
	} `json:"request"`
	SQLfunc func(e *DataInternalMainCreateUser)
}
type DataInternalMainGetUser struct {
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
	SQLfunc func(e *DataInternalMainGetUser)
}
type DataInternalMainGetUserPassHash struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserLogin
	} `json:"request"`
	Response struct {
		DataInternalMainUserPassHash
	} `json:"response"`
	SQLfunc func(e *DataInternalMainGetUserPassHash)
}
type DataInternalMainUpdateUser struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
		DataInternalMainUserName
		DataInternalMainUserAvatar
	} `json:"request"`
	SQLfunc func(e *DataInternalMainUpdateUser)
}
type DataInternalMainDeleteUser struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
	} `json:"request"`
	SQLfunc func(e *DataInternalMainDeleteUser)
}
type DataInternalMainUpdateUserPassHash struct {
	Package
	DataInternalMain
	Request struct {
		DataInternalMainUserLogin
		DataInternalMainUserPassHash
	} `json:"request"`
	SQLfunc func(e *DataInternalMainUpdateUserPassHash)
}

func (e *DataInternalMainCreateUser) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainGetUser) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainGetUserPassHash) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainUpdateUser) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainDeleteUser) SQL() {
	e.SQLfunc(e)
}
func (e *DataInternalMainUpdateUserPassHash) SQL() {
	e.SQLfunc(e)
}

func (e *DataInternalMainCreateUser) StatusPtr() *Status {
	return &e.Status
}
func (e *DataInternalMainGetUser) StatusPtr() *Status {
	return &e.Status
}
func (e *DataInternalMainGetUserPassHash) StatusPtr() *Status {
	return &e.Status
}
func (e *DataInternalMainUpdateUser) StatusPtr() *Status {
	return &e.Status
}
func (e *DataInternalMainDeleteUser) StatusPtr() *Status {
	return &e.Status
}
func (e *DataInternalMainUpdateUserPassHash) StatusPtr() *Status {
	return &e.Status
}

// Data Errors
type DataInternalServiceError struct{ Status }
type DataExternalServiceError struct{ Status }
type DataIncorrectArgumentError struct{ Status }
type DataNotFoundError struct{ Status }
type DataAlreadyExistError struct{ Status }
type DataAccessDenied struct{ Status }
