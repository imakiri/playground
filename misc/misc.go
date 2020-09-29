package misc

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
