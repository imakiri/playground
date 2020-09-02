package a

import (
	"fmt"
	"github.com/imakiri/playground/goroutines/inter"
)

type Foo bool

func (Foo) Do(str string) string {
	fmt.Printf("a.Do %s\n", str)
	return str
}

type Bar bool

func (Bar) Does(f inter.Foo) {
	fmt.Printf("a.Does\n")
	f.Do("Does")
}
