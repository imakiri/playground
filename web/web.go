package web

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/protos"
	"net/http"
)

var r = mux.NewRouter()
var rr = &http.ServeMux{}
var s = &http.Server{}
var sr = &http.Server{}
var gc protos.FaceDetecterClient

type Execute interface {
	SQL()
}

func NewWebServer(gcNew protos.FaceDetecterClient) error {
	gc = gcNew

	rr.HandleFunc("/", redirect)
	sr.Handler = rr

	go func() {
		_ = sr.ListenAndServe()
	}()

	RegisterHandlers(r)
	s.Handler = r
	return s.ListenAndServeTLS("cert.pem", "privkey.pem")
}

func RegisterHandlers(rr *mux.Router) {
	rr.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	rr.Handle("/", GetRoot{})
	rr.Handle("/detect", GetRootDetect{})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}
