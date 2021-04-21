package asset

import (
	"context"
	"github.com/imakiri/gorum/internal/transport"
	"github.com/imakiri/gorum/internal/types"
	"io/ioutil"
)

type Service struct {
	transport.UnimplementedAssetsServer
	assets *types.Assets
}

func (s *Service) Get(_ context.Context, _ *types.Request) (*types.Assets, error) {
	var ass types.Assets
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

func (s *Service) Reload(_ context.Context, _ *types.Request) error {
	var assets types.Assets
	var err error

	if assets.Index, err = ioutil.ReadFile("assets/index.html"); err != nil {
		return err
	}
	if assets.CSS, err = ioutil.ReadFile("assets/style.css"); err != nil {
		return err
	}
	if assets.Ico, err = ioutil.ReadFile("assets/ico.png"); err != nil {
		return err
	}

	s.assets = &assets
	return err
}
