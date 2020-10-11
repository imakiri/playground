package inside

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// General SQL execute method
type EXEC interface {
	ExecuteSQL() error
}

// MAIN Data Entity
type MAIN struct {
	db *sqlx.DB
}

// MAIN_User Fields
type MAIN_User_Id struct {
	Id uint `db:"id"`
}
type MAIN_User_Login struct {
	Login string `db:"login"`
}
type MAIN_User_Avatar struct {
	Avatar []byte `db:"avatar"`
}
type MAIN_User_Name struct {
	Name string `db:"name"`
}
type MAIN_User_PassHash struct {
	PassHash []byte `db:"passHash"`
}

// MAIN_Data Fields
type MAIN_Data_Id struct {
	Id uint `db:"id"`
}
type MAIN_Data_Date struct {
	Date time.Time `db:"date"`
}
type MAIN_Data_Query struct {
	Query string `db:"query"`
}
type MAIN_Data_Pic struct {
	Pic []byte `db:"pic"`
}
type MAIN_Data_UserId struct {
	UserId uint `db:"userId"`
}
