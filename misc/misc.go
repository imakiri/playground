package misc

import (
	"fmt"
	"github.com/imakiri/playground/data/inside"
	"net/http"
	"reflect"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}

}

func Test8() {
	var c = inside.MAIN_GetUser_1{
		MAIN: inside.Main,
		EXEC: nil,
		Request: struct {
			inside.MAIN_User_Id
			inside.MAIN_User_Login
		}{},
		Response: nil,
	}

	c.Request.Login = "imakiri"

	if err := c.ExecuteSQL(); err != nil {
		fmt.Print(err.Error() + "\n")
		fmt.Print(reflect.TypeOf(err))
	} else {
		fmt.Print(c.Response)
	}
}
