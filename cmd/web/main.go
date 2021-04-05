package main

import (
	"flag"
	"fmt"
	"github.com/imakiri/gorum/web"
	"log"
)

const path_cert = "secrets/grpc/cert.crt"

type opts struct {
	debug  bool
	domain string
	port   string
}

func NewLauncher(o opts) (*Launcher, error) {
	var l Launcher
	var err error
	l.debug = o.debug
	l.statusWeb = make(chan error)
	l.statusRedirector = make(chan error)

	if l.debug {
		l.web, err = web.NewService(l.statusWeb, false)
		if err != nil {
			return nil, err
		}

		return &l, err
	}

	l.redirector, err = web.NewHTTPSRedirector(l.statusRedirector)
	if err != nil {
		return nil, err
	}

	l.web, err = web.NewService(l.statusWeb, true)
	if err != nil {
		return nil, err
	}

	return &l, err
}

type Launcher struct {
	debug            bool
	web              *web.Service
	redirector       *web.HTTPSRedirector
	statusWeb        chan error
	statusRedirector chan error
}

func (l *Launcher) Launch() error {
	if l.debug {
		l.web.Launch()
		return <-l.statusWeb
	} else {
		var err error
		l.web.Launch()
		l.redirector.Launch()

		select {
		case err = <-l.statusRedirector:
			l.web.Stop()
		case err = <-l.statusWeb:
			l.redirector.Stop()
		}

		return err
	}
}

const (
	domain       = "imakiri-ips.ddns.net"
	default_port = "25565"
)

func main() {
	var o opts

	var debug = flag.Bool("debug", true, "set to false to launch a production ready system")
	flag.Parse()

	o.debug = *debug
	o.domain = domain
	o.port = default_port

	fmt.Println(o)

	var l, err = NewLauncher(o)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(l.Launch())
}
