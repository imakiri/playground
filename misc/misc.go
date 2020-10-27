package misc

import (
	"fmt"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}

}

func Test8() {
	c := data.NewRequest(data.RequestInternalMainGetUser{}).(*core.DataInternalMainGetUser)

	c.Request.Login = "imakiri"
	c.SQL()

	if c.Package.Status.IsOK() {
		fmt.Print(c.Response)
		fmt.Print("\nOK")
	} else {
		fmt.Print(c.Package.Status.Error())
	}
}
