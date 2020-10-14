package core

import (
	"github.com/jmoiron/sqlx"
	"time"
)

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
