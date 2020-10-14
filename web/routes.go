package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

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
		e := core.NewError(core.WebTemplateParseError{}, err.Error())
		fmt.Print(e.Error())
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		e := core.NewError(core.WebTemplateExecuteError{}, err.Error())
		fmt.Print(e.Error())
	}

	fmt.Printf("%v WebGetRoot passed to %s\n", time.Now(), r.RemoteAddr)
}
func (e GetRootAssetsCss_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	f, err := ioutil.ReadFile("web/templates/assets/css/style.css")
	if err != nil {
		e := core.NewError(core.WebCssReadError{}, err.Error())
		fmt.Print(e.Error())
	}

	_, _ = w.Write(f)
}
func (e GetRootIco_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")

	_, _ = w.Write(icoF)
}
func (e GetRootUserLogin_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var c = data.NewRequest(data.RequestInternalMainGetUser_1{}).(*core.DataInternalMainGetUser_1)

	c.Request.Login = vars["login"]
	c.SQL()

	if err := c.Package.Error; err != nil {
		_, _ = fmt.Fprintf(w, "%s [%s]", c.Package.Error, c.Request.Login)
	} else {
		_, _ = fmt.Fprintf(w, "%s", c.Response.Name)
	}
}
