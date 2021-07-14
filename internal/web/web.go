package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/asset/transport"
	"github.com/imakiri/gorum/internal/utils"
	"html/template"
	"net/http"
)

func ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)

	_, _ = w.Write([]byte(err.Error()))
}

func push(p http.Pusher) error {
	var err error
	if err = p.Push("/assets/css", nil); err != nil {
		return err
	}
	if err = p.Push("/assets/ico", nil); err != nil {
		return err
	}

	return err
}

type webService struct {
	serviceAsset transport.AssetClient
	assets       *transport.Assets
	template     struct {
		home  *template.Template
		gorum *template.Template
	}
	router *mux.Router
}

func (s *webService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func (s *webService) load() error {
	var assets, err = s.serviceAsset.Get(context.Background(), &transport.Request{})
	if err != nil {
		return err
	}
	s.assets = assets

	s.template.home = template.New("root")
	if s.template.home, err = s.template.home.Parse(string(s.assets.Index)); err != nil {
		return err
	}
	if s.template.home, err = s.template.home.Parse(string(s.assets.Home)); err != nil {
		return err
	}

	s.template.gorum = template.New("gorum")
	if s.template.gorum, err = s.template.gorum.Parse(string(s.assets.Index)); err != nil {
		return err
	}
	if s.template.gorum, err = s.template.gorum.Parse(string(s.assets.Gorum)); err != nil {
		return err
	}

	return nil
}

func NewWebService(asset transport.AssetClient) (*webService, error) {
	if utils.IsNil(asset) {
		return nil, erres.NilArgument
	}

	var s = new(webService)
	s.router = mux.NewRouter()
	s.serviceAsset = asset

	var err = s.load()
	if err != nil {
		return nil, err
	}

	s.router.HandleFunc("/", s.root)
	s.router.HandleFunc("/assets/css", s.rootAssetsCss)
	s.router.HandleFunc("/assets/ico", s.rootAssetsIco)
	s.router.HandleFunc("/gorum", s.rootGorum)

	return s, nil
}
