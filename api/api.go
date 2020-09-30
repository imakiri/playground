package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Resolve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}

func Run(rr *mux.Router) error {
	return nil
}
