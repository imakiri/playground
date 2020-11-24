package web

import (
	"fmt"
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/core"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type GetRoot_1 struct{}
type GetRootUserLogin_1 struct{}
type GetRootDetect struct {
	memCache [][]byte
}

// Web ServeHTTP Methods

// GET /
func (e GetRoot_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("%v WebGetRoot passed to %s\n", time.Now(), r.RemoteAddr)
}

//// GET /user/{login}
//func (e GetRootUserLogin_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	var c = data.NewRequest(data.RequestInternalMainGetUser{}).(*core.DataInternalMainGetUser)
//
//	c.Request.Login = vars["login"]
//	Execute.SQL(c)
//
//	if err := c.Package.Error; err != nil {
//		_, _ = fmt.Fprintf(w, "%s [%s]", c.Package.Error(), c.Request.Login)
//	} else {
//		_, _ = fmt.Fprintf(w, "%s", c.Response.Name)
//	}
//}

func (e GetRootDetect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("File Upload Endpoint Hit")
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

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	reImg, err := app.Detect(fileBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")

	//id, err := strconv.Atoi(mux.Vars(r)["id"])
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//}
	_, _ = w.Write(reImg)
}
