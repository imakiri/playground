package web

import (
	"context"
	"fmt"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/protos"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type GetRoot struct{}
type PostRootDetect struct{}

// Web ServeHTTP Methods

// GET /
func (e GetRoot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("/assets/css/style.css", nil)
		_ = p.Push("/assets/favicon.ico", nil)
	}
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		e := core.NewStatus(core.WebTemplateParseError{}, err)
		fmt.Print(e.Error())
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		e := core.NewStatus(core.WebTemplateExecuteError{}, err)
		fmt.Print(e.Error())
	}

	fmt.Printf("%v WebGetRoot enpoint hit and pass to %s\n", time.Now(), r.RemoteAddr)
}

// POST /detect
func (e PostRootDetect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("PostRootDetect endpoint hit")
	_ = r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	response, _ := gc.Detect(context.Background(), &protos.DetectionRequest{Img: fileBytes})
	if err := response.GetErr(); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.String()))
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	_, _ = w.Write(response.GetImg().GetData())
}
