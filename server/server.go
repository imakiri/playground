package server

import (
	"github.com/gorilla/mux"
	_ "github.com/imakiri/playground/api"
	"github.com/imakiri/playground/web"
	_ "github.com/imakiri/playground/web"
	"log"
	"net/http"
)

var r = mux.NewRouter()
var rr = &http.ServeMux{}
var s = &http.Server{}
var sr = &http.Server{}

func Run() {
	rr.HandleFunc("/", redirect)
	sr.Handler = rr

	go func() {
		_ = sr.ListenAndServe()
	}()

	_ = web.Run(r)
	s.Handler = r
	log.Fatal(s.ListenAndServeTLS("C:/Certbot/live/imakiri.ddns.net/cert.pem", "C:/Certbot/live/imakiri.ddns.net/privkey.pem"))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}
