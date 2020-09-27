package app

import (
	"github.com/imakiri/playground/core"
	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) (re core.Re) {
	re.Data, re.Err = bcrypt.GenerateFromPassword([]byte(pass+salt), hashCost)
	return
}
