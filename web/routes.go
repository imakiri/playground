package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"html/template"
	"net/http"
	"time"
)

type GetRoot_1 struct{}
type GetRootUserLogin_1 struct{}

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

// GET /user/{login}
func (e GetRootUserLogin_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var c = data.NewRequest(data.RequestInternalMainGetUser{}).(*core.DataInternalMainGetUser)

	c.Request.Login = vars["login"]
	Execute.SQL(c)

	if err := c.Package.Error; err != nil {
		_, _ = fmt.Fprintf(w, "%s [%s]", c.Package.Error(), c.Request.Login)
	} else {
		_, _ = fmt.Fprintf(w, "%s", c.Response.Name)
	}
}
