package data

import (
	"github.com/imakiri/playground/data/internal"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
	"time"
)

type external struct {
	client *http.Client
	key    string
}

var f []byte

var Internal = internal.Consecutive
var InternalC = internal.Concurrent
var External external

func Init() (err error) {
	defer func() { f = nil }()

	err = internal.Init()
	if err != nil {
		return
	}

	External.client = &http.Client{Timeout: time.Second}
	f, err = ioutil.ReadFile("data/key")
	if err != nil {
		return
	}

	External.key = string(f)
	return
}

func GetSalt() string {
	return internal.Salt
}

func RunTest() {}
