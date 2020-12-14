package web

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/core"
	"net/http"
)

var Web core.Web

// Returns prepared router and redirectionRouter for http web server, or error
func NewWebRouters(s core.Settings) (*mux.Router, *mux.Router, error) {
	Web = s.Web
	a, err := app.NewApp(s)
	if err != nil {
		return nil, nil, err
	}

	var router = mux.NewRouter()
	var redirRouter = mux.NewRouter()

	redirRouter.HandleFunc("/", redirect)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	router.Handle("/", root{a})
	router.Handle("/detect", detect{a})

	return router, redirRouter, nil
}

// Redirect from HTTP to HTTPS
func redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}

// Internal Service Error Response
func ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
