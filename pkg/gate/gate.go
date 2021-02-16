package gate

import (
	"github.com/imakiri/playground/core"
)

type Gate struct {
}

func NewGate(s core.Settings) (*Gate, error) {
	var g Gate

	return &g, nil
}
