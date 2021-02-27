package web

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func (e *Service) Root(w http.ResponseWriter, r *http.Request) {
	n := time.Now().Format(time.RFC822)
	fmt.Printf("%s Root web-endpoint hit by %s\n", n, r.RemoteAddr)

	//if p, ok := w.(http.Pusher); ok {
	//	_ = p.Push("assets/css/style.css", nil)
	//	_ = p.Push("assets/favicon.ico", nil)
	//}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("frontend/assets/html/index.html")
	if err != nil {
		ise(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		ise(w, err)
		return
	}
}

//type detect struct {
//	app *app2.App
//}
//
//
//type root struct {
//	app *app2.App
//}
//
//// GET /
//func (e root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if p, ok := w.(http.Pusher); ok {
//		_ = p.Push("/assets/css/style.css", nil)
//		_ = p.Push("/assets/favicon.ico", nil)
//	}
//	w.Header().Set("Content-Type", "text/html")
//
//	t, err := template.ParseFiles("web/templates/index.html")
//	if err != nil {
//		ui.ise(w, err)
//		return
//	}
//
//	err = t.ExecuteTemplate(w, "index", nil)
//	if err != nil {
//		ui.ise(w, err)
//		return
//	}
//
//	fmt.Printf("%v WebGetRoot endpoint hit by %s\n", time.Now(), r.RemoteAddr)
//}
//
//// POST /detect
//func (e detect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	u := ui.readUUIDFromCookie(w, r)
//	_ = r.ParseMultipartForm(10 << 20)
//
//	file, _, err := r.FormFile("file")
//	if err != nil {
//		ui.ise(w, err)
//		return
//	}
//
//	fileBytes, err := ioutil.ReadAll(file)
//	if err != nil {
//		ui.ise(w, err)
//		return
//	}
//
//	img, err := e.app.Detect(u, fileBytes)
//	if err != nil {
//		ui.ise(w, err)
//		return
//	}
//
//	fmt.Printf("%v WebPostDetect endpoint hit by %s\n", time.Now(), r.RemoteAddr)
//
//	w.Header().Set("Content-Type", "image/jpeg")
//	_, _ = w.Write(img)
//}
