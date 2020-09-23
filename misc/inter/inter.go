package inter

type Foo interface {
	Do(str string) string
}

type Bar interface {
	Does(f Foo)
}
