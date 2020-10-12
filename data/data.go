package data

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
)

var Salt string

func init() {
	f, err := ioutil.ReadFile("data/salt")
	if err != nil {
		panic(err)
	}

	Salt = string(f)
}
