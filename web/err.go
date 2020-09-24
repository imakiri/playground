package web

import (
	"fmt"
	"log"
)

type checkImp struct{}

func (checkImp) CSS(err error) {
	if err != nil {
		log.Fatalf("Web/CSS error/%s\n", err)
	}
}

func (checkImp) ICO(err error) {
	if err != nil {
		log.Fatalf("Web/Init/ICO error/%s\n", err)
	}
	fmt.Print("Web/Init/ICO passed\n")
}
