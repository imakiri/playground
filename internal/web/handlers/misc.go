package handlers

import (
	"net/http"
)

func Go2HTTPS(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}

func ISE(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)

	_, _ = w.Write([]byte(err.Error()))
}

func push(p http.Pusher) error {
	var err error
	if err = p.Push("/assets/css", nil); err != nil {
		return err
	}
	if err = p.Push("/assets/ico", nil); err != nil {
		return err
	}

	return err
}
