package app

import (
	"github.com/imakiri/playground/server/core"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

func Hash(pass string, wg *sync.WaitGroup, c chan core.Re) {
	defer wg.Done()

	var re core.Re
	re.Data, re.Err = bcrypt.GenerateFromPassword([]byte(pass+salt), hashCost)
	c <- re
}
