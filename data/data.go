package data

import (
	"github.com/imakiri/playground/data/external"
	"github.com/imakiri/playground/data/internal"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Internal = internal.Release
var External = external.Release

func Init() (err error) {
	err = internal.Init()
	if err != nil {
		return
	}

	err = external.Init()
	if err != nil {
		return
	}

	return
}

func GetSalt() string {
	return internal.Salt
}

func RunTest() {}
