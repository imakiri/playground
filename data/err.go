package data

import "fmt"

type InitError struct {
	typ string
	err string
}

func (e InitError) Error() string {
	return fmt.Sprintf("%s error: %s", e.typ, e.err)
}
