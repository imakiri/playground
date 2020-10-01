package data

import (
	"github.com/imakiri/playground/data/inside"
	"github.com/imakiri/playground/data/outside"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var External = outside.Release

func Init() (err error) {
	err = inside.Init()
	if err != nil {
		return &InitError{"Internal", err.Error()}
	}

	err = outside.Init()
	if err != nil {
		return &InitError{"External", err.Error()}
	}

	return
}

func GetSalt() string {
	return inside.Salt
}

func Internal() *inside.R {
	return &inside.Release
}

func RunTest() {}
