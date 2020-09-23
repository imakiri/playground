package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	id      uint
	login   string
	avatar  []byte
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

type Re struct {
	User
	Data
	Err error
}

var f []byte
var main *sql.DB
var ro *sql.Rows
var re Re
var err error
var c checkImp
var query q
var salt string

func init() {
	f, err = ioutil.ReadFile("dsn")
	check.DSN(c, err)

	main, err = sql.Open("mysql", string(f))
	check.Connection(c, err)

	err = main.Ping()
	check.Ping(c, err)

	f, err = ioutil.ReadFile("salt")
	check.Salt(c, err)
	salt = string(f)

	f, err = ioutil.ReadFile("server/app/local/sql/getPassHash.sql")
	check.HashQ(c, err)
	query.getPassHash = string(f)

	f, err = ioutil.ReadFile("server/app/local/sql/getProfile.sql")
	check.ProfileQ(c, err)
	query.getProfile = string(f)
}

func Run() {
	re := User{}
	ro, err = main.Query(query.getProfile)
	checkImp.Query(c, err)
	defer func() {
		ro.Close()
	}()

	ro.Next()
	_ = ro.Scan(&re.id, &re.name, &re.encPass)
	fmt.Printf("User/id: %v, name: %s, encPass: %b", re.id, re.name, re.encPass)
}

func GetSalt() string {
	return salt
}

func GetUserProfile(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	const qGetProfile = "SELECT name, avatar FROM main.users WHERE login = ?"
	ro, err = main.Query(qGetProfile, login)
	if re.Check(err, c) {
		return
	}

	ro.Next()
	err = ro.Scan(&re.User.name, &re.User.avatar)
	if re.Check(err, c) {
		return
	}

	c <- re
}

//func SetProfile(login string, uData User, wg *sync.WaitGroup, c chan Re) {
//	defer wg.Done()
//
//
//	ro, err = main.Query(, login)
//	if re.Check(err, c) {
//		return
//	}
//
//	c <- re
//}

// Login procedure
func GetPassHash(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	const qGetPassHash = "SELECT passHash FROM main.users WHERE name = ?"
	ro, err = main.Query(qGetPassHash, login)
	if re.Check(err, c) {
		return
	}

	ro.Next()
	err = ro.Scan(&re.Data)
	if re.Check(err, c) {
		return
	}

	c <- re
}
