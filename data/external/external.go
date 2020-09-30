package external

import (
	"io/ioutil"
	"net/http"
	"time"
)

type r struct{}

var f []byte
var client *http.Client
var key string

var Release r

func Init() (err error) {
	client = &http.Client{Timeout: 5 * time.Second}

	f, err = ioutil.ReadFile("data/key")
	if err != nil {
		return
	}

	key = string(f)
	f = nil
	return
}
