package server

import (
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
)

var s *http2.Server
var l net.Listener
var err error
var c checkImp
var conn net.Conn

type check interface {
	Http2(err error)
	Conn(err error)
}

func Run2() {
	l, err = net.Listen("tcp", "0.0.0.0:1010")
	check.Http2(c, err)

	for {
		conn, err = l.Accept()
		check.Conn(c, err)
		s.ServeConn(conn, &http2.ServeConnOpts{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
		})})
	}
}
