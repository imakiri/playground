package schema

import "time"

type User struct {
	Id       uint   `db:"id"`
	Login    string `db:"login"`
	Avatar   []byte `db:"avatar"`
	Name     string `db:"name"`
	PassHash []byte `db:"passHash"`
}

type Data struct {
	Id     uint      `db:"id"`
	Date   time.Time `db:"date"`
	Query  string    `db:"query"`
	Pic    []byte    `db:"pic"`
	UserId uint      `db:"userId"`
}

type Re struct {
	User
	Data
	Err error
}
