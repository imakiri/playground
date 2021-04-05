package web

import (
	"github.com/imakiri/gorum/web/internal"
	"net/http"
)

func (s *Service) css(w http.ResponseWriter, r *http.Request) {
	_ = internal.CSS(s.data.css, w, r)
}

func (s *Service) ico(w http.ResponseWriter, r *http.Request) {
	_ = internal.ICO(s.data.ico, w, r)
}

func (s *Service) root(w http.ResponseWriter, r *http.Request) {
	_ = internal.Root(w, r)
}
