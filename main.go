package main

import (
	"github.com/imakiri/playground/server"
	_ "github.com/imakiri/playground/server/store"
)

func main() {
	//store.Run()

	server.Run()
}
