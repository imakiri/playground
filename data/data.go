package data

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type check interface {
	DSN(err error)
	Connection(err error)
	Ping(err error)
	Salt(err error)
	Key(err error)
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

type internal struct {
	main *sqlx.DB
	salt string
}

type external struct {
	client *http.Client
	key    string
}

var f []byte
var re Re
var err error
var c checkImp
var Internal internal
var External external

func init() {
	f, err = ioutil.ReadFile("data/dsn")
	check.DSN(c, err)

	Internal.main, err = sqlx.Open("mysql", string(f))
	check.Connection(c, err)

	err = Internal.main.Ping()
	check.Ping(c, err)

	f, err = ioutil.ReadFile("data/salt")
	check.Salt(c, err)

	Internal.salt = string(f)
	External.client = &http.Client{Timeout: time.Second}

	f, err = ioutil.ReadFile("data/key")
	check.Key(c, err)

	External.key = string(f)
}

func RunTest() {
	wg := sync.WaitGroup{}
	c := make(chan Re, 1)

	wg.Add(1)
	Internal.GetUserPassHash("imakiri", &wg, c)
	wg.Wait()

	res := <-c

	fmt.Printf("PassHash: %s", res.PassHash)
}
