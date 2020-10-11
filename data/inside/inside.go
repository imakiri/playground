package inside

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

var goquDB *goqu.Database
var sqlxDB *sqlx.DB
var main *sql.DB
var Salt string
var Main MAIN

func init() {
	f, err := ioutil.ReadFile("data/dsn")
	if err != nil {
		panic(err)
	}

	main, err = sql.Open("mysql", string(f))
	if err != nil {
		panic(err)
	}

	sqlxDB, err = sqlx.Open("mysql", string(f))
	if err != nil {
		panic(err)
	}

	err = main.Ping()
	if err != nil {
		panic(err)
	}

	f, err = ioutil.ReadFile("data/salt")
	if err != nil {
		panic(err)
	}

	Salt = string(f)
	goquDB = goqu.Dialect("mysql").DB(main)

	Main.db = sqlxDB
}
