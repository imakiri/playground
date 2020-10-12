package data

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

var Connection_Internal_Main Internal_Main

func init() {
	f, err := ioutil.ReadFile("data/dsn")
	if err != nil {
		panic(err)
	}

	Connection_Internal_Main.SQLX_DB, err = sqlx.Open("mysql", string(f))
	if err != nil {
		panic(err)
	}

	err = Connection_Internal_Main.SQLX_DB.Ping()
	if err != nil {
		panic(err)
	}
}
