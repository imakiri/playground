package local

import (
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
var ERR error

func init() {
	DB, ERR = gorm.Open("sqlite3", "server/storage/storage.db")
	if ERR != nil {
		log.Fatal(ERR)
	}
}
