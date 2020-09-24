package data

import (
	"fmt"
	"log"
)

type checkImp struct{}

func (checkImp) DSN(err error) {
	if err != nil {
		log.Fatalf("Internal/Init/DSN error/%s\n", err)
	}
	fmt.Print("Internal/Init/DSN passed\n")
}

func (checkImp) Connection(err error) {
	if err != nil {
		log.Fatalf("Internal/Init/Connection error/%s\n", err)
	}
	fmt.Print("Internal/Init/Connection passed\n")
}

func (checkImp) Ping(err error) {
	if err != nil {
		log.Fatalf("Internal/Init/Ping error/%s\n", err)
	}
	fmt.Print("Internal/Init/Ping passed\n")
}

func (checkImp) Salt(err error) {
	if err != nil {
		log.Fatalf("Internal/Init/Salt error/%s\n", err)
	}
	fmt.Print("Internal/Init/Salt passed\n")
}

func (checkImp) Key(err error) {
	if err != nil {
		log.Fatalf("External/Init/Key error/%s\n", err)
	}
	fmt.Print("External/Init/Key passed\n")
}
