package server

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/api"
	_ "github.com/imakiri/playground/api"
	"github.com/imakiri/playground/web"
	_ "github.com/imakiri/playground/web"
	"log"
	"net/http"
)

var r = mux.NewRouter()

func RunS() {
	var s = &http.Server{}

	_ = api.Run(r)
	_ = web.Run(r)
	s.Handler = r

	log.Fatal(s.ListenAndServeTLS("C:/Certbot/live/imakiri.ddns.net/cert.pem", "C:/Certbot/live/imakiri.ddns.net/privkey.pem"))
}
