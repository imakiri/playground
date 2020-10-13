package web

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
)

var icoF []byte

func init() {
	f, err := ioutil.ReadFile("web/templates/favicon.ico")
	if err != nil {
		log.Fatal(ERROR_InitIco{ERROR(err.Error())})
	}

	icoF = f
}

func RegisterHandlers(rr *mux.Router) {
	rr.Handle("/", GetRoot_1{})
	rr.Handle("/assets/css/style.css", GetRootAssetsCss_1{})
	rr.Handle("/favicon.ico", GetRootIco_1{})
	rr.Handle("/user/{user}", GetRootUserLogin_1{})
}
