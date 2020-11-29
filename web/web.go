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

	go func() {
		_ = redirServer.ListenAndServe()
	}()

	RegisterHandlers(router)
	server.Handler = router
	return server.ListenAndServeTLS("cert.pem", "privkey.pem")
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
