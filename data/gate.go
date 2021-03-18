package data

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/erres"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type ConfigGate interface {
	Get4DataGate(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.DataGate, error)
}

type Gate struct {
	configCached *cfg.DataGate
	config       ConfigGate
	db           *sqlx.DB
}

func NewGate(cg ConfigGate) (*Gate, error) {
	var s Gate
	var err error

	s.config = cg
	s.configCached, err = s.config.Get4DataGate(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	s.db, err = sqlx.Connect("pgx", s.configCached.GetDSN())
	if err != nil {
		return nil, erres.E_ConnectionError.SetTime("").SetDescription(err.Error())
	}

	return &s, err
}
