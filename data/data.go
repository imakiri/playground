package data

import (
	"github.com/imakiri/playground/data/inside"
	"github.com/imakiri/playground/data/outside"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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
	return &inside.R{}
}

func External() *outside.R {
	return &outside.R{}
}
