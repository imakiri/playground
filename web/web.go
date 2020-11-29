package web

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/protos"
	"net/http"
)

var router = mux.NewRouter()
var redirRouter = &http.ServeMux{}
var server = &http.Server{}
var redirServer = &http.Server{}
var gc protos.FaceDetecterClient

func NewWebServer(gcNew protos.FaceDetecterClient) error {
	gc = gcNew

	redirRouter.HandleFunc("/", redirect)
	redirServer.Handler = redirRouter
	RegisterHandlers(router)
	server.Handler = router

	rsc := make(chan error)
	sc := make(chan error)

	go func(rsc chan error) {
		rsc <- redirServer.ListenAndServe()
	}(rsc)

	go func(sc chan error) {
		sc <- server.ListenAndServeTLS("cert.pem", "privkey.pem")
	}(sc)

	select {
	case err := <-rsc:
		return err
	case err := <-sc:
		return err
	}
}

func RegisterHandlers(rr *mux.Router) {
	rr.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	rr.Handle("/", GetRoot{})
	rr.Handle("/detect", PostRootDetect{})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}
