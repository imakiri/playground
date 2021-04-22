package asset

import (
	"context"
	"github.com/imakiri/gorum/internal/asset/transport"
	"io/ioutil"
)

type Service struct {
	transport.UnimplementedAssetServer
	assets *transport.Assets
}

func (s *Service) Get(_ context.Context, _ *transport.Request) (*transport.Assets, error) {
	var ass transport.Assets
	var err error

	if ass.Index, err = ioutil.ReadFile("assets/index.html"); err != nil {
		return nil, err
	}
	if ass.CSS, err = ioutil.ReadFile("assets/style.css"); err != nil {
		return nil, err
	}
	if ass.Ico, err = ioutil.ReadFile("assets/ico.png"); err != nil {
		return nil, err
	}

	return &ass, err
}
