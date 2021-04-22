package handlers

import (
	"net/http"
)

func CSS(css []byte, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_, _ = w.Write(css)
}

func Ico(ico []byte, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_, _ = w.Write(ico)
}
