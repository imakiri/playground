package misc

import (
	"fmt"
	data "github.com/imakiri/playground/data"
	"net/http"
	"reflect"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}

}

func Test8() {
	var c = data.Internal_Main_Method_GetUser_1{
		Internal_Main: data.Connection_Internal_Main,
		Request: struct {
			data.Internal_Main_User_Id
			data.Internal_Main_User_Login
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
