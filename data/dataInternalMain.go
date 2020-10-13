package data

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// General SQL execute method
type Execute interface {
	SQL() error
}

// InternalMain Data Entity
type InternalMain struct {
	SQLX_DB *sqlx.DB
}

// InternalMainUser Fields
type InternalMainUserId struct {
	Id uint `db:"id"`
}
type InternalMainUserLogin struct {
	Login string `db:"login"`
}
type InternalMainUserAvatar struct {
	Avatar []byte `db:"avatar"`
}
type InternalMainUserName struct {
	Name string `db:"name"`
}
type InternalMainUserPassHash struct {
	PassHash []byte `db:"passHash"`
}

// InternalMainData Fields
type InternalMainDataId struct {
	Id uint `db:"id"`
}
type InternalMainDataDate struct {
	Date time.Time `db:"date"`
}
type InternalMainDataQuery struct {
	Query string `db:"query"`
}
type InternalMainDataPic struct {
	Pic []byte `db:"pic"`
}
type InternalMainDataUserId struct {
	UserId uint `db:"userId"`
}
