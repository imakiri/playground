package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Thing struct {
	Header string
	Data   []byte
	Error  error
}

func (th *Thing) GetHeader() string {
	return th.Header
}

func (th *Thing) GetData() []byte {
	return th.Data
}

func (th *Thing) GetError() error {
	return th.Error
}

type local struct {
	isInitialized bool
}

func (l local) GetThing(str string, c chan Thing) {

}

func (l local) DoThing(str string, c chan Thing) {

}

func (l local) ChangeThing(str string, th Thing, c chan Thing) {

}

func (l local) StoreThing(str string, th Thing, c chan Thing) {

}

var Local local

var db *gorm.DB

var err error

func init() {
	db, err = gorm.Open("sqlite3", "server/storage.db")
	if err != nil {
		log.Fatal(err)
	}
	Local.isInitialized = true
}

func Run() {}
