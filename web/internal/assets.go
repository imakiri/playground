package internal

import (
	"net/http"
)

func ICO(ico []byte, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "image/png")
	var _, err = w.Write(ico)
	return err
}

func CSS(css []byte, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/css")
	var _, err = w.Write(css)
	return err
}
