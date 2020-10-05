package app

import (
	"fmt"
	"github.com/imakiri/playground/data"
	"github.com/imakiri/playground/data/schema"
	_ "golang.org/x/crypto/bcrypt"
	"log"
)

type Re struct {
}

const hashCost = 10

var salt string
var reData schema.Main

func Init() (err error) {
	if err := data.Init(); err != nil {
		log.Fatal(err.Error())
	}

	salt = data.GetSalt()
	return
}

func RunTest1() {
	re, err := IsAuthorized("imakiri", "imakiri")
	fmt.Printf("Main: %v, %v\n", re, err)
}

func IsAuthorized(login string, pass string) (bool, error) {
	//defer func() {
	//	reData = schema.Main{}
	//}()
	//
	//reData.User.Login = login
	//
	//switch err := data.Internal().GetUserPassHash(&reData.User).(type) {
	//default:
	//	reData.Err = err
	//}
	//
	//if reData.Err != nil {
	//	return false, reData.Err
	//}
	//
	//if bcrypt.CompareHashAndPassword(reData.PassHash, []byte(pass)) == nil {
	//	return true, nil
	//} else {
	//	return false, nil
	//}
	return false, nil
}
