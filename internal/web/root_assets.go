package web

import "net/http"

func (s *webService) rootAssetsCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_, _ = w.Write(s.assets.Css)
}

func (s *webService) rootAssetsIco(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public")
	w.Header().Set("Cache-Control", "max-age=86400")
	_, _ = w.Write(s.assets.Ico)
}
