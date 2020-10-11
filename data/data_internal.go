package data

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

var Salt string
var Connection_Internal_Main Internal_Main

func init() {
	f, err := ioutil.ReadFile("data/dsn")
	if err != nil {
		panic(err)
	}

	Connection_Internal_Main.Db, err = sqlx.Open("mysql", string(f))
	if err != nil {
		panic(err)
	}

	err = Connection_Internal_Main.Db.Ping()
	if err != nil {
		panic(err)
	}

	f, err = ioutil.ReadFile("data/salt")
	if err != nil {
		panic(err)
	}

	Salt = string(f)
}
