package main

import (
	"fmt"
	"log"
	"net/http"
)

var err error
var port int

func main() {
	_, _ = fmt.Print("Enter port: ")
	_, err = fmt.Scanf("%d", &port)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := http.Server{Addr: fmt.Sprintf(":%d", port)}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "Кибер-сычевальня Мага полупопия")
	})

	_, _ = fmt.Printf("\nServer is running on port %d", port)
	log.Fatal(server.ListenAndServe())
	//
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
//	db.GetPassHash("imakiri", &wg, c)
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
