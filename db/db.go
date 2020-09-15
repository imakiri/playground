package db

import (
	"database/sql"
	"fmt"
	"github.com/imakiri/playground/server/core"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"sync"
	"time"
)

var main *sql.DB
var ro *sql.Rows
var err error
var c checkImp
var salt string

type check interface {
	DSN(err error)
	Connection(err error)
	Ping(err error)
	Query(err error)
	Salt(err error)
	Hash(err error)
}

type User struct {
	id      uint
	name    string
	encPass []byte
}

type Data struct {
	id     uint
	date   time.Time
	query  string
	pic    []byte
	userId uint
}

func init() {
	var f []byte
	f, err = ioutil.ReadFile("dsn")
	check.DSN(c, err)

	main, err = sql.Open("mysql", string(f))
	check.Connection(c, err)

	err = main.Ping()
	check.Ping(c, err)

	f, err = ioutil.ReadFile("salt")
	check.Salt(c, err)
	salt = string(f)
}

func Run() {
	re := User{}
	ro, err = main.Query("SELECT id, name, encPass FROM main.users")
	checkImp.Query(c, err)
	defer func() {
		ro.Close()
	}()

	ro.Next()
	_ = ro.Scan(&re.id, &re.name, &re.encPass)
	fmt.Printf("User/id: %v, name: %s, encPass: %b", re.id, re.name, re.encPass)
}

func GetPassHash(name string, wg *sync.WaitGroup, c chan core.Re) {
	defer wg.Done()

	var re core.Re
	var ro *sql.Rows
	ro, err = main.Query("SELECT encPass FROM main.users where name = ?", name)
	if re.Check(err, c) {
		return
	}

	err = ro.Scan(&re.Data)
	if re.Check(err, c) {
		return
	}

	c <- re
}
