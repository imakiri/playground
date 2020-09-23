package server

import (
	"github.com/gorilla/mux"
	_ "github.com/imakiri/playground/server/api"
	"github.com/imakiri/playground/server/web"
	_ "github.com/imakiri/playground/server/web"
	"log"
	"net/http"
)

func RunS() {
	//_ = api.Run(Router)

	var Router = mux.NewRouter()
	var Server = http.Server{}

	_ = web.Run(Router)
	Server.Handler = Router

	log.Fatal(Server.ListenAndServeTLS("C:/Certbot/live/imakiri.ddns.net/cert.pem", "C:/Certbot/live/imakiri.ddns.net/privkey.pem"))
}
