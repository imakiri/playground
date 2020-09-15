package app

import (
	"fmt"
	"log"
)

type checkImp struct{}

func (checkImp) Salt(err error) {
	if err != nil {
		log.Fatalf("App/Salt error/%s\n", err)
	}
	fmt.Print("App/Salt passed\n")
}
