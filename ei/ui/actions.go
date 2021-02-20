package ui

import (
	"fmt"
	app2 "github.com/imakiri/playground/pkg/app"
	"io/ioutil"
	"net/http"
	"time"
)

type detect struct {
	app *app2.App
}

// POST /detect
func (e detect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := readUUIDFromCookie(w, r)
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

	img, err := e.app.Detect(u, fileBytes)
	if err != nil {
		ise(w, err)
		return
	}

	fmt.Printf("%v WebPostDetect endpoint hit by %s\n", time.Now(), r.RemoteAddr)

	w.Header().Set("Content-Type", "image/jpeg")
	_, _ = w.Write(img)
}
