package inside

import (
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

type R struct{}

var f []byte
var main *sqlx.DB
var Salt string

var Release R

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
	f = nil
	return
}

type InternalServiceError struct {
	err string
}

func (e InternalServiceError) Error() string {
	return e.err
}

type IncorrectArgumentError struct {
	err string
}

func (e IncorrectArgumentError) Error() string {
	return e.err
}

type NotFoundError struct {
	err string
}

func (e NotFoundError) Error() string {
	return e.err
}

type UserAlreadyExistError struct {
	err string
}

func (e UserAlreadyExistError) Error() string {
	return e.err
}
