package internal

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

type consecutive struct{}
type concurrent struct{}
type erroneous struct{}

var f []byte
var main *sqlx.DB
var Salt string

var Consecutive consecutive
var Concurrent concurrent
var Erroneous erroneous

func Init() (err error) {
	f, err = ioutil.ReadFile("data/dsn")
	if err != nil {
		return
	}

	main, err = sqlx.Open("mysql", string(f))
	if err != nil {
		return
	}

	err = main.Ping()
	if err != nil {
		return
	}

	f, err = ioutil.ReadFile("data/salt")
	if err != nil {
		return
	}

	Salt = string(f)
	return
}
