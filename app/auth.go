package app

import (
	"github.com/imakiri/playground/data/schema"
	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) (re schema.Re) {
	re.User.PassHash, re.Err = bcrypt.GenerateFromPassword([]byte(pass+salt), hashCost)
	return
}
