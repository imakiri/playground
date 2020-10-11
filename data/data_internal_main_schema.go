package data

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// General SQL execute method
type Exec interface {
	ExecuteSQL() error
}

// Internal_Main Data Entity
type Internal_Main struct {
	Db *sqlx.DB
}

// MAIN_User Fields
type Internal_Main_User_Id struct {
	Id uint `db:"id"`
}
type Internal_Main_User_Login struct {
	Login string `db:"login"`
}
type Internal_Main_User_Avatar struct {
	Avatar []byte `db:"avatar"`
}
type Internal_Main_User_Name struct {
	Name string `db:"name"`
}
type Internal_Main_User_PassHash struct {
	PassHash []byte `db:"passHash"`
}

// MAIN_Data Fields
type Internal_Main_Data_Id struct {
	Id uint `db:"id"`
}
type Internal_Main_Data_Date struct {
	Date time.Time `db:"date"`
}
type Internal_Main_Data_Query struct {
	Query string `db:"query"`
}
type Internal_Main_Data_Pic struct {
	Pic []byte `db:"pic"`
}
type Internal_Main_Data_UserId struct {
	UserId uint `db:"userId"`
}
