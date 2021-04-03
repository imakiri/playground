package internal

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegistrarMain(s *http.Server) error {
	var err error
	var router = mux.NewRouter()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", root)
	router.HandleFunc("/forum", rootForum)

	s.Handler = router
	return err
}
