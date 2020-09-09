package server

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api"
	"github.com/imakiri/playground/server/web"
	"log"
	"net/http"
)

var Router = mux.NewRouter().Schemes("http").Subrouter()

var Server = &http.Server{
	Addr: ":80",
}

func Run() {
	_ = api.Run(Router)
	_ = web.Run(Router)

	Server.Handler = Router

	log.Fatal(Server.ListenAndServe())
}
