package goroutines

import (
	"github.com/imakiri/playground/goroutines/a"
	"github.com/imakiri/playground/goroutines/inter"
)

var N a.Foo

var D a.Bar

func Vu() {
	inter.Bar.Does(D, N)
}
