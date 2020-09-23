package main

import (
	"github.com/imakiri/playground/db"
	"net/http"
)

func main() {
	db.RunTest()
	//t1()
	//server.Run()
	//goroutines.Vu()

}

//func t1() {
//	//db.Run()
//
//	var re1 core.Re
//	var re2 core.Re
//	wg := sync.WaitGroup{}
//	c := make(chan core.Re, 2)
//
//	wg.Add(1)
//	app.Hash("неверныйпароль", &wg, c)
//
//	wg.Add(1)
//	db.GetUserPassHash("imakiri", &wg, c)
//
//	re1 = <- c
//	re2 = <- c
//
//	fmt.Printf("%s\n", re1.Data)
//	fmt.Printf("%s\n", re2.Data)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}
}
