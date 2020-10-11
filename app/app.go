package app

import (
	"fmt"
	"github.com/imakiri/playground/data"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

const hashCost = 10

func RunTest1() {
	err := IsAuthorized("imakiri", "imakiri")
	fmt.Printf("Main: %v\n", err)
}

func IsAuthorized(login string, pass string) (err error) {
	c := data.Internal_Main_Method_GetUserPassHash_1{
		Internal_Main: data.Connection_Internal_Main,
		Request: struct {
			data.Internal_Main_User_Login
		}{},
		Response: struct {
			data.Internal_Main_User_PassHash
		}{},
	}

	c.Request.Login = login

	switch e := c.ExecuteSQL().(type) {
	case error:
		return e
	}

	if bcrypt.CompareHashAndPassword(c.Response.PassHash, []byte(pass)) == nil {
		return nil
	} else {
		return ERROR_NotAuthorized{}
	}
}
