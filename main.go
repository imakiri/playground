package main

import (
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/server"
	"log"
)

func main() {
	server.Run()
	//test()
}

func test() {
	err := app.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 4; i > 0; i-- {
		app.RunTest1()
	}
}
