package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"sync"
	"time"
)

type check interface {
	DSN(err error)
	Connection(err error)
	Ping(err error)
	Query(err error)
	Salt(err error)
	Hash(err error)
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
var salt string

func init() {
	f, err = ioutil.ReadFile("db/dsn")
	check.DSN(c, err)

	main, err = sqlx.Open("mysql", string(f))
	check.Connection(c, err)

	err = main.Ping()
	check.Ping(c, err)

	f, err = ioutil.ReadFile("db/salt")
	check.Salt(c, err)
	salt = string(f)
}

func RunTest() {
	wg := sync.WaitGroup{}
	c := make(chan Re, 1)

	wg.Add(1)
	GetUserPassHash("imakiri", &wg, c)
	wg.Wait()

	res := <-c

	fmt.Printf("PassHash: %s", res.PassHash)
}

func GetSalt() string {
	return salt
}

func GetUser(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE login = ?", login)
	if re.Check(err, c) {
		return
	}

	c <- re
	close(c)
}

func GetUserById(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = main.Get(&re.User, "SELECT name, avatar FROM main.users WHERE id = ?", id)
	if re.Check(err, c) {
		return
	}

	c <- re
	close(c)
}

func GetUserPassHash(login string, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	err = main.Get(&re.User, "SELECT passHash FROM main.users WHERE login = ?", login)
	if re.Check(err, c) {
		return
	}

	c <- re
	close(c)
}

func GetData(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//err = main.Get(&re.User, "SELECT pic FROM main.data WHERE userId = ?", id)
	//if re.Check(err, c) {
	//	return
	//}

	c <- re
	close(c)
}

func AddPic(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//err = main.Get(&re.User, "SELECT pic FROM main.data WHERE userId = ?", id)
	//if re.Check(err, c) {
	//	return
	//}

	c <- re
	close(c)
}

func CreateUser(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//err = main.Get(&re.User, "SELECT pic FROM main.data WHERE userId = ?", id)
	//if re.Check(err, c) {
	//	return
	//}

	c <- re
	close(c)
}

func UpdateUser(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//err = main.Get(&re.User, "SELECT pic FROM main.data WHERE userId = ?", id)
	//if re.Check(err, c) {
	//	return
	//}

	c <- re
	close(c)
}

func DeleteUser(id uint, wg *sync.WaitGroup, c chan Re) {
	defer wg.Done()

	//err = main.Get(&re.User, "SELECT pic FROM main.data WHERE userId = ?", id)
	//if re.Check(err, c) {
	//	return
	//}

	c <- re
	close(c)
}
