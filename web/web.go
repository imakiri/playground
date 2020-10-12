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
	rr.Handle("/", GET_Root{})
	rr.Handle("/assets/css/style.css", GET_Root_Assets_CSS_1{})
	rr.Handle("/favicon.ico", GET_Root_Ico_1{})
}
