package web

import (
	"fmt"
	"github.com/imakiri/playground/app"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type root struct {
	app *app.App
}

// GET /
func (e root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("/assets/css/style.css", nil)
		_ = p.Push("/assets/favicon.ico", nil)
	}
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		ise(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		ise(w, err)
		return
	}

	fmt.Printf("%v WebGetRoot enpoint hit by %s\n", time.Now(), r.RemoteAddr)
}

type detect struct {
	app *app.App
}

// POST /detect
func (e detect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")
	if err != nil {
		ise(w, err)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ise(w, err)
		return
	}

	img, err := e.app.Detect(fileBytes)
	if err != nil {
		ise(w, err)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	_, _ = w.Write(img)
}
