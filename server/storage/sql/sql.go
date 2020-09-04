package sql

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

var err error

func init() {
	db, err = gorm.Open("sqlite3", "server/storage.db")
	if err != nil {
		log.Fatal(err)
	}
}
