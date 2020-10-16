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
	PassHash []byte `db:"passHash" json:"dataInternalMainUserPassHash"`
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
	UserId uint `db:"userId" json:"dataInternalMainDataUserId"`
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
