package asset

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/gorum/internal/utils"
	"net/http"
)

func Handler(s *Service) *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/css", s.css)
	router.HandleFunc("/ico", s.ico)
	return router
}

func (s *Service) css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_ = utils.SendBytes(s.assets.CSS, w, r)
}

func (s *Service) ico(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_ = utils.SendBytes(s.assets.Ico, w, r)
}
