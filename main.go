package main

import (
	"github.com/imakiri/playground/server"
	_ "github.com/imakiri/playground/server/storage"
)

func main() {
	//storage.Run()

	server.Run()
	//goroutines.Vu()
}
