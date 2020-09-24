package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/web/models"
	"io/ioutil"
	"net/http"
)

type check interface {
	CSS(err error)
	ICO(err error)
}

var f []byte
var err error
var c checkImp
var icoF []byte

func init() {
	f, err = ioutil.ReadFile("web/templates/favicon.ico")
	check.ICO(c, err)
	icoF = f
}

func test(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Assets")
}

func Run(rr *mux.Router) error {
	rr.Handle("/", models.Root)
	rr.HandleFunc("/assets/css/style.css", css)
	rr.HandleFunc("/favicon.ico", ico)
	return nil
}

func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	f, err = ioutil.ReadFile("web/templates/assets/css/style.css")
	check.CSS(c, err)

	_, _ = w.Write(f)
}

func ico(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")

	_, _ = w.Write(icoF)
}
