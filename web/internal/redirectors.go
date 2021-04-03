package internal

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RedirectorHTTPS(s *http.Server) error {
	var err error
	var router = mux.NewRouter()

	router.HandleFunc("/", go2https)

	s.Handler = router
	return err
}
