package app

import (
	"github.com/imakiri/playground/data"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

const hashCost = 10

// Returns error if it is not authorized
func CheckAuthorization(login string, pass string) (err error) {
	e := data.InternalMainGetUserPassHash_1{
		InternalMain: data.ConnectionInternalMain,
		Request: struct {
			data.InternalMainUserLogin
		}{},
		Response: struct {
			data.InternalMainUserPassHash
		}{},
	}

	e.Request.Login = login

	switch e := e.SQL().(type) {
	case error:
		return e
	}

	if bcrypt.CompareHashAndPassword(e.Response.PassHash, []byte(pass)) == nil {
		return nil
	} else {
		return ERROR_NotAuthorized{}
	}
}

func Img() {

}
