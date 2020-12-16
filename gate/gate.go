package gate

import (
	"github.com/google/uuid"
	"github.com/imakiri/playground/core"
)

type Gate struct {
}

func (e *Gate) GetIdFromUUID(u uuid.UUID) (*core.Identity, error) {

	return nil, nil
}

func (e *Gate) CheckPermissionsOf(id *core.Identity) core.Checker {
	return &Checker{id, nil}
}

type Checker struct {
	id  *core.Identity
	err error
}

func (e *Checker) For(funcName core.FuncName) (bool, error) {
	return true, nil
}

func NewGate(s core.Settings) (*Gate, error) {
	var g Gate

	return &g, nil
}
