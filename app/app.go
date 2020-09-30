package app

import (
	"fmt"
	"github.com/imakiri/playground/data"
	"github.com/imakiri/playground/data/schema"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Re struct {
}

const hashCost = 10

var salt string
var reData schema.Re

func Init() (err error) {
	if err := data.Init(); err != nil {
		log.Fatal(err.Error())
	}

	salt = data.GetSalt()
	return
}

func RunTest1() {
	re, err := IsAuthorized("imakiri", "imakiri")
	fmt.Printf("Re: %v, %v\n", re, err)
}

func IsAuthorized(login string, pass string) (bool, error) {
	defer func() {
		reData = schema.Re{}
	}()

	reData = data.Internal.GetUserPassHash(login)
	if reData.Err != nil {
		return false, reData.Err
	}

	if bcrypt.CompareHashAndPassword(reData.PassHash, []byte(pass)) == nil {
		return true, nil
	} else {
		return false, nil
	}
}
