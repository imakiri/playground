package web

import "net/http"

func (s *webService) rootStaticCss(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_, _ = w.Write(s.content.Static.Css)
}

func (s *webService) rootStaticIco(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_, _ = w.Write(s.content.Static.Ico)
}
