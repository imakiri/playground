package db

import (
	"fmt"
	"log"
)

type checkImp struct{}

func (checkImp) DSN(err error) {
	if err != nil {
		log.Fatalf("DB/DSN error/%s\n", err)
	}
	fmt.Print("DB/DSN passed\n")
}

func (checkImp) Connection(err error) {
	if err != nil {
		log.Fatalf("DB/Connection error/%s\n", err)
	}
	fmt.Print("DB/Connection initialized\n")
}

func (checkImp) Ping(err error) {
	if err != nil {
		log.Fatalf("DB/Ping error/%s\n", err)
	}
	fmt.Print("DB/Ping passed\n")
}

func (checkImp) Query(err error) {
	if err != nil {
		log.Fatalf("DB/Query error/%s\n", err)
	}
	fmt.Print("DB/Query passed\n")
}

func (checkImp) Salt(err error) {
	if err != nil {
		log.Fatalf("App/Salt error/%s\n", err)
	}
	fmt.Print("App/Salt passed\n")
}

func (checkImp) Hash(err error) {
	if err != nil {
		log.Fatalf("App/Hash error/%s\n", err)
	}
	fmt.Print("App/Hash passed\n")
}
