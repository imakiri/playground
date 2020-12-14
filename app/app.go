package app

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
)

type App struct {
	app      core.App
	services *core.Services
	data     *data.DB
}

func NewApp(s core.Settings) (*App, error) {
	var a App
	var err error

	a.app = s.App
	a.data, err = data.NewDB(s)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (e *App) DB() *data.DB {
	return e.data
}
