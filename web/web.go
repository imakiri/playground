package web

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/core"
	"io/ioutil"
	"log"
)

type Execute interface {
	SQL() error
}

var icoF []byte

func init() {
	f, err := ioutil.ReadFile("web/templates/favicon.ico")
	if err != nil {
		log.Fatal(core.NewError(core.WebIcoInitError{}, err.Error()))
	}

	icoF = f
}

func RegisterHandlers(rr *mux.Router) {
	rr.Handle("/", GetRoot_1{})
	rr.Handle("/assets/css/style.css", GetRootAssetsCss_1{})
	rr.Handle("/favicon.ico", GetRootIco_1{})
	rr.Handle("/user/{login}", GetRootUserLogin_1{})
}
