package web

import "net/http"

func ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)

	_, _ = w.Write([]byte(err.Error()))
}

func push(p http.Pusher) error {
	var err error
	if err = p.Push("/static/css", nil); err != nil {
		return err
	}
	if err = p.Push("/static/ico", nil); err != nil {
		return err
	}
	return nil
}
