package app

const hashCost = 0

var salt string
var err error
var c checkImp

type check interface {
	Salt(err error)
}

func init() {}
