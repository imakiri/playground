package server

import "log"

type checkImp struct{}

func (checkImp) Http2(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (checkImp) Conn(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
