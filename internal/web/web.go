package web

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/web/content"
	"github.com/imakiri/gorum/internal/web/transport"
	"github.com/imakiri/gorum/pkg/utils"
	"html/template"
	"io/ioutil"
	"net/http"
)

const path = "internal/web/content/"

type contentService struct {
	transport.UnimplementedContentServer
}

func (s *contentService) Get(context.Context, *transport.Request) (*transport.Content, error) {
	var c = new(transport.Content)
	var err error

	if c.Main, err = ioutil.ReadFile(path + "main.html"); err != nil {
		return nil, err
	}
	if c.Index, err = ioutil.ReadFile(path + "index.html"); err != nil {
		return nil, err
	}
	if c.StaticCss, err = ioutil.ReadFile(path + "static/style.css"); err != nil {
		return nil, err
	}
	if c.StaticIco, err = ioutil.ReadFile(path + "static/ico.png"); err != nil {
		return nil, err
	}
	if c.GorumMain, err = ioutil.ReadFile(path + "gorum/main.html"); err != nil {
		return nil, err
	}
	if c.GorumIndex, err = ioutil.ReadFile(path + "gorum/index.html"); err != nil {
		return nil, err
	}

	fmt.Println("content/get")
	return c, nil
}

func NewContentService() (*contentService, error) {
	var s = new(contentService)

	var _, err = s.Get(context.Background(), new(transport.Request))
	if err != nil {
		return nil, err
	}

	return s, nil
}

type webService struct {
	debug    bool
	services struct {
		content transport.ContentClient
	}
	content *content.Content
	router  *mux.Router
}

func (s *webService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if s.debug {
		var err = s.load()
		if err != nil {
			fmt.Println(err)
		}
	}
	s.router.ServeHTTP(writer, request)
}

func (s *webService) load() error {
	var raw, err = s.services.content.Get(context.Background(), &transport.Request{})
	if err != nil {
		return err
	}

	s.content = new(content.Content)
	s.content.Static.Css = raw.StaticCss
	s.content.Static.Ico = raw.StaticIco
	s.content.Index = template.New("index")

	if s.content.Index, err = s.content.Index.Parse(string(raw.Main)); err != nil {
		return err
	}
	if s.content.Index, err = s.content.Index.Parse(string(raw.Index)); err != nil {
		return err
	}

	s.content.Gorum.Index = template.New("index")
	if s.content.Gorum.Index, err = s.content.Gorum.Index.Parse(string(raw.Main)); err != nil {
		return err
	}
	if s.content.Gorum.Index, err = s.content.Gorum.Index.Parse(string(raw.GorumMain)); err != nil {
		return err
	}
	if s.content.Gorum.Index, err = s.content.Gorum.Index.Parse(string(raw.GorumIndex)); err != nil {
		return err
	}

	return nil
}

func NewWebService(debug bool, contentClient transport.ContentClient) (*webService, error) {
	if utils.IsNil(contentClient) {
		return nil, erres.NilArgument
	}

	var s = new(webService)
	s.debug = debug
	s.router = mux.NewRouter()
	s.services.content = contentClient

	var err = s.load()
	if err != nil {
		return nil, err
	}

	s.router.HandleFunc("/", s.root)
	s.router.HandleFunc("/static/css", s.rootStaticCss)
	s.router.HandleFunc("/static/ico", s.rootStaticIco)
	s.router.HandleFunc("/gorum", s.rootGorum)

	return s, nil
}
