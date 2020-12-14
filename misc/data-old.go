package misc

import (
	"github.com/jmoiron/sqlx"
	"time"
)

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

type SQLContainer interface {
	ExecuteSQL() error
	DB() *sqlx.DB
}

// DataInternalMain Data Entities
type DataInternalMainCreateUser struct {
	db      *sqlx.DB
	Request struct {
		DataInternalMainUserLogin
		DataInternalMainUserAvatar
		DataInternalMainUserName
		DataInternalMainUserPassHash
	} `json:"request"`
	SQLfunc func(e *DataInternalMainCreateUser) error
}
type DataInternalMainGetUser struct {
	db      *sqlx.DB
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
	} `json:"request"`
	Response struct {
		DataInternalMainUserAvatar
		DataInternalMainUserName
	} `json:"response"`
	SQLfunc func(e *DataInternalMainGetUser) error
}
type DataInternalMainGetUserPassHash struct {
	db      *sqlx.DB
	Request struct {
		DataInternalMainUserLogin
	} `json:"request"`
	Response struct {
		DataInternalMainUserPassHash
	} `json:"response"`
	SQLfunc func(e *DataInternalMainGetUserPassHash) error
}
type DataInternalMainUpdateUser struct {
	db      *sqlx.DB
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
		DataInternalMainUserName
		DataInternalMainUserAvatar
	} `json:"request"`
	SQLfunc func(e *DataInternalMainUpdateUser) error
}
type DataInternalMainDeleteUser struct {
	db      *sqlx.DB
	Request struct {
		DataInternalMainUserId
		DataInternalMainUserLogin
	} `json:"request"`
	SQLfunc func(e *DataInternalMainDeleteUser) error
}
type DataInternalMainUpdateUserPassHash struct {
	db      *sqlx.DB
	Request struct {
		DataInternalMainUserLogin
		DataInternalMainUserPassHash
	} `json:"request"`
	SQLfunc func(e *DataInternalMainUpdateUserPassHash) error
}

func (e *DataInternalMainCreateUser) ExecuteSQL() error {
	return e.SQLfunc(e)
}
func (e *DataInternalMainGetUser) ExecuteSQL() error {
	return e.SQLfunc(e)
}
func (e *DataInternalMainGetUserPassHash) ExecuteSQL() error {
	return e.SQLfunc(e)
}
func (e *DataInternalMainUpdateUser) ExecuteSQL() error {
	return e.SQLfunc(e)
}
func (e *DataInternalMainDeleteUser) ExecuteSQL() error {
	return e.SQLfunc(e)
}
func (e *DataInternalMainUpdateUserPassHash) ExecuteSQL() error {
	return e.SQLfunc(e)
}

func (e *DataInternalMainCreateUser) DB() *sqlx.DB {
	return e.db
}
func (e *DataInternalMainGetUser) DB() *sqlx.DB {
	return e.db
}
func (e *DataInternalMainGetUserPassHash) DB() *sqlx.DB {
	return e.db
}
func (e *DataInternalMainUpdateUser) DB() *sqlx.DB {
	return e.db
}
func (e *DataInternalMainDeleteUser) DB() *sqlx.DB {
	return e.db
}
func (e *DataInternalMainUpdateUserPassHash) DB() *sqlx.DB {
	return e.db
}
