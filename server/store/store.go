package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/mail"
	"time"
)

var db *gorm.DB

var err error

func init() {
	db, err = gorm.Open("sqlite3", "server/store.db")
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {}

type Data struct {
	User     user
	Location location
	Visit    visit
}

type person struct {
	name     [2]string
	gender   uint8
	birthday time.Time
}

type user struct {
	id     uint
	email  mail.Address
	person person
}

type country struct {
	country uint8
}

type site struct {
	id          uint
	name        string
	country     *country
	distance    float32
	description string
}

type visit struct {
	id   uint16
	site *site
	user *user
	date time.Time
	mark int8
}
