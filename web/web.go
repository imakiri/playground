package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Execute interface {
	SQL()
}

func init() {}

func RegisterHandlers(rr *mux.Router) {
	rr.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	rr.Handle("/", GetRoot_1{})
	rr.Handle("/user/{login}", GetRootUserLogin_1{})
}
