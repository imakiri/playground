package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/data"
	"html/template"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

// Web Handlers
type GetRoot_1 struct{}
type GetRootAssetsCss_1 struct{}
type GetRootIco_1 struct{}
type GetRootUserLogin_1 struct{}

// Web ServeHTTP Methods
func (e GetRoot_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("assets/css/style.css", nil)
	}
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		e := ERROR_ParseTemplate{ERROR(err.Error())}
		fmt.Printf("Error ocured: %s, %s\n", reflect.TypeOf(e), e.Error())
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		e := ERROR_ExecuteTemplate{ERROR(err.Error())}
		fmt.Printf("Error ocured: %s, %s\n", reflect.TypeOf(e), e.Error())
	}

	fmt.Printf("%v Web/Root passed to %s\n", time.Now(), r.RemoteAddr)
}
func (e GetRootAssetsCss_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	f, err := ioutil.ReadFile("web/templates/assets/css/style.css")
	if err != nil {
		e := ERROR_ReadCss{ERROR(err.Error())}
		fmt.Printf("Error ocured: %s, %s\n", reflect.TypeOf(e), e.Error())
	}

	_, _ = w.Write(f)
}
func (e GetRootIco_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")

	_, _ = w.Write(icoF)
}
func (e GetRootUserLogin_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var c = data.InternalMainGetUser_1{
		InternalMain: data.ConnectionInternalMain,
		Request: struct {
			data.InternalMainUserId
			data.InternalMainUserLogin
		}{},
		Response: struct {
			data.InternalMainUserAvatar
			data.InternalMainUserName
		}{},
	}

	c.Request.Login = vars["user"]

	switch err := data.Execute.SQL(&c).(type) {
	case error:
		_, _ = fmt.Fprint(w, err.Error()+"\n")
	default:
		_, _ = fmt.Print(c.Response.Name + "\n")
		_, _ = fmt.Fprint(w, c.Response.Name+"\n")
	}
}
