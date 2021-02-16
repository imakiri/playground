package app

import (
	"github.com/google/uuid"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/pkg/data"
	"github.com/imakiri/playground/pkg/gate"
)

type App struct {
	app  core.App
	data Data
	gate Gate
}

type Data interface {
	CreateUser(c *core.ContainerCreateUser) error
	GetPassHash(c *core.ContainerGetUserPassHash) error
}

type Gate interface {
	GetIdFromUUID(u uuid.UUID) (*core.Identity, error)
	CheckPermissionsOf(id *core.Identity) core.Checker
}

func checkPermissionForUUID(g Gate, u uuid.UUID, fn core.FuncName) bool {
	id, err := g.GetIdFromUUID(u)
	if err != nil {
		return false
	}

	b, err := g.CheckPermissionsOf(id).For(fn)
	if err != nil || !b {
		return false
	} else {
		return true
	}
}

func NewApp(s core.Settings) (*App, error) {
	var a App
	var err error

	a.data, err = data.NewDB(s)
	if err != nil {
		return nil, err
	}

	g, err := gate.NewGate(s)
	if err != nil {
		return nil, err
	}

	a.gate = g
	a.app = s.App

	return &a, nil
}
