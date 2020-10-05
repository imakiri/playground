package inside

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

type R struct{}

var goquDB *goqu.Database
var main *sql.DB
var Salt string

var Release R

func init() {
	f, err := ioutil.ReadFile("data/dsn")
	if err != nil {
		panic(err)
	}

	main, err = sql.Open("mysql", string(f))
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
}

type BaseError string

func (b BaseError) Error() string {
	return string(b)
}

type InternalServiceError struct {
	BaseError
}

type IncorrectArgumentError struct {
	BaseError
}

type NotFoundError struct {
	BaseError
}

type UserAlreadyExistError struct {
	BaseError
}

type NoUserToDelete struct {
	BaseError
}
