package asset

import (
	"context"
	"fmt"
	"github.com/imakiri/gorum/internal/asset/transport"
	"io/ioutil"
)

type Service struct {
	transport.UnimplementedAssetServer
	assets *transport.Assets
}

func (s *Service) Get(_ context.Context, _ *transport.Request) (*transport.Assets, error) {
	var ass = new(transport.Assets)
	var err error

	if ass.Index, err = ioutil.ReadFile("assets/index.html"); err != nil {
		return nil, err
	}
	if ass.Css, err = ioutil.ReadFile("assets/style.css"); err != nil {
		return nil, err
	}
	if ass.Ico, err = ioutil.ReadFile("assets/ico.png"); err != nil {
		return nil, err
	}
	if ass.Home, err = ioutil.ReadFile("assets/home.html"); err != nil {
		return nil, err
	}
	if ass.Gorum, err = ioutil.ReadFile("assets/gorum.html"); err != nil {
		return nil, err
	}

	fmt.Println("Assets.Get")

	return ass, err
}

func NewService() (*Service, error) {
	var s = new(Service)
	var _, err = s.Get(context.Background(), new(transport.Request))
	if err != nil {
		return nil, err
	}

	return s, nil
}
