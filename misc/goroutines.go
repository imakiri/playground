package misc

import (
	"fmt"
	"github.com/imakiri/playground/misc/a"
	"github.com/imakiri/playground/misc/inter"
)

var N a.Foo

var D a.Bar

func Vu() {
	inter.Bar.Does(D, N)
}

func GoroutTest() {
	c := make(chan int, 5)

	for i := 5; i > 0; i-- {
		go sf(10-i, i, c)
	}

	var re []int
	for i := 5; i > 0; i-- {
		re = append(re, <-c)
	}

	fmt.Printf("%v", re)
}

func sf(a, b int, c chan int) {
	c <- 2*a + b
}
