package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/asset/transport"
	internalHttp "github.com/imakiri/gorum/internal/http"
	"github.com/imakiri/gorum/internal/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net/http"
)

func connect(domain string, port string) (*grpc.ClientConn, error) {
	var err error
	var ca []byte

	if ca, err = ioutil.ReadFile("secrets/ca/ca.crt"); err != nil {
		return nil, err
	}

	var cp = x509.NewCertPool()
	if !cp.AppendCertsFromPEM(ca) {
		return nil, erres.Error("certificate error").Extend(0)
	}

	var conf = &tls.Config{
		RootCAs:            cp,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(domain+":"+port, grpc.WithTransportCredentials(credentials.NewTLS(conf)))
	if err != nil {
		return nil, err
	}

	return conn, err
}

func NewLauncher(debug bool, domain string, port string) (*Launcher, error) {
	var l Launcher
	l.debug = debug
	l.statusWeb = make(chan error)
	l.statusRedirector = make(chan error)

	var cc, err = connect(domain, port)
	if err != nil {
		return nil, err
	}

	var ass = transport.NewAssetClient(cc)

	var ws http.Handler
	ws, err = web.NewWebService(ass)
	if err != nil {
		return nil, err
	}

	if l.debug {
		l.web, err = internalHttp.NewServer(ws, l.statusWeb, false)
		if err != nil {
			return nil, err
		}

		return &l, err
	}

	l.redirector, err = internalHttp.NewRedirector(l.statusRedirector)
	if err != nil {
		return nil, err
	}

	l.web, err = internalHttp.NewServer(ws, l.statusWeb, true)
	if err != nil {
		return nil, err
	}

	return &l, err
}

type Launcher struct {
	debug            bool
	asset            transport.AssetClient
	web              *internalHttp.Server
	redirector       *internalHttp.Redirector
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
	domain = "imakiri-ips.ddns.net"
	port   = "25565"
)

func main() {
	var debug = flag.Bool("debug", true, "set to false to launch a production ready system")
	flag.Parse()

	var l, err = NewLauncher(*debug, domain, port)
	if err != nil {
		log.Fatalln(err)
	}

	err = l.Launch()
	log.Fatalln(err)
}
