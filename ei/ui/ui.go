package ui

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/pkg/app"
	"net/http"
)

var Web core.UI

// Returns prepared router and redirectionRouter for http web server, or error
func NewWebRouters(s core.Settings) (*mux.Router, *mux.Router, error) {
	Web = s.UI
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

func readUUIDFromCookie(w http.ResponseWriter, r *http.Request) uuid.UUID {
	var u uuid.UUID
	c, err := r.Cookie("uuid")
	if err != nil {
		u = uuid.New()
		http.SetCookie(w, &http.Cookie{Name: "uuid", Value: u.String()})
	} else {
		u, err = uuid.Parse(c.Value)
		if err != nil {
			u = uuid.New()
			http.SetCookie(w, &http.Cookie{Name: "uuid", Value: u.String()})
		}
	}
	return u
}
