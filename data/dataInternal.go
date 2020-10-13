package data

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

var ConnectionInternalMain InternalMain

func init() {
	f, err := ioutil.ReadFile("data/dsn")
	if err != nil {
		panic(err)
	}

	ConnectionInternalMain.SQLX_DB, err = sqlx.Open("mysql", string(f))
	if err != nil {
		panic(err)
	}

	err = ConnectionInternalMain.SQLX_DB.Ping()
	if err != nil {
		panic(err)
	}
}
