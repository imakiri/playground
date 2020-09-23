package misc

import (
	"github.com/imakiri/playground/misc/a"
	"github.com/imakiri/playground/misc/inter"
)

var N a.Foo

var D a.Bar

func Vu() {
	inter.Bar.Does(D, N)
}
