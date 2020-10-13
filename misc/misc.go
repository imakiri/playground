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
	var e = data.InternalMainGetUser_1{
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

	e.Request.Login = "imakiri"

	if err := e.SQL(); err != nil {
		fmt.Print(err.Error() + "\n")
		fmt.Print(reflect.TypeOf(err))
	} else {
		fmt.Print(e.Response)
	}
}
