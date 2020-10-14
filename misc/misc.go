package misc

import (
	"fmt"
	"github.com/imakiri/playground/core"
	data "github.com/imakiri/playground/data"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}

}

func Test8() {
	c := data.NewRequest(data.RequestInternalMainGetUser_1{}).(*core.DataInternalMainGetUser_1)

	c.Request.Login = "imakiri"
	c.SQL()

	if err := c.Package.Error; err != nil {
		fmt.Print(err.Error() + "\n")
	} else {
		fmt.Print(c.Response)
	}
}
