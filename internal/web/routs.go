package web

import (
	"fmt"
	"github.com/imakiri/gorum/internal/utils"
	"github.com/imakiri/gorum/internal/web/handlers"
	"net/http"
	"time"
)

func (s *Service) load(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = s.Load(); err != nil {
		handlers.ISE(w, err)
	}
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

func (s *Service) root(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] web.root hit by ip:%s\n", t, r.RemoteAddr)

	w.Header().Set("Content-Type", "text/html")

	var err = handlers.Root(s.index, w, r)
	if err != nil {
		handlers.ISE(w, err)
	}
}
