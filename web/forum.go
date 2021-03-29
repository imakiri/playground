package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerForum struct {
	router *mux.Router
}

func (h HandlerForum) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func newHandlerForum() (*HandlerForum, error) {
	var h HandlerForum
	var err error

	h.router = mux.NewRouter()
	return &h, err
}
