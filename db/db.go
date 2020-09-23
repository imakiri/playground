package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"sync"
	"time"
)

type q struct {
	getPassHash string
	getProfile  string
}

type check interface {
	DSN(err error)
	Connection(err error)
	Ping(err error)
	Query(err error)
	Salt(err error)
	Hash(err error)
	HashQ(err error)
	ProfileQ(err error)
}

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

var f []byte
var main *sqlx.DB
var re Re
var err error
var c checkImp
var query q
var salt string

func init() {
	f, err = ioutil.ReadFile("dsn")
	check.DSN(c, err)

	main, err = sqlx.Open("mysql", string(f))
	check.Connection(c, err)

	err = main.Ping()
	check.Ping(c, err)

	f, err = ioutil.ReadFile("salt")
	check.Salt(c, err)
	salt = string(f)
}

func RunTest() {
	wg := sync.WaitGroup{}
	c := make(chan Re, 1)

	wg.Add(1)
	GetPassHash("imakiri", &wg, c)
	wg.Wait()

	res := <-c
	fmt.Printf("PassHash: %s", res.PassHash)
}

func GetSalt() string {
	return salt
}

func GetUserProfile(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE login = ?", login)
	if re.Check(err, c) {
		return
	}

	c <- re
}

func GetPassHash(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = main.Get(&re.User, "SELECT passHash FROM main.users WHERE name = ?", login)
	if re.Check(err, c) {
		return
	}

	c <- re
}
