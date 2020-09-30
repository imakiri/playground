package misc

import "strings"

type Gyto struct {
	lik int
	lpe string
}

func (g *Gyto) Lik() *int {
	return &g.lik
}

func (g *Gyto) Lpe() *string {
	return &g.lpe
}

func Uuy(in string) (out string) {
	out = strings.ToLower(in)
	return
}
