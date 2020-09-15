package app

import "io/ioutil"

const hashCost = 0

var salt string
var err error
var c checkImp

type check interface {
	Salt(err error)
}

func init() {
	var f []byte
	f, err = ioutil.ReadFile("salt")
	check.Salt(c, err)
	salt = string(f)
}
