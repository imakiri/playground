package remote

import (
	"github.com/imakiri/playground/core"
	"net"
	"net/http"
	"sync"
	"time"
)

var netTransport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:  5 * time.Second,
		Deadline: time.Now().Add(5 * time.Second),
	}).DialContext,
	TLSHandshakeTimeout: 5 * time.Second,
}
var Client = &http.Client{
	Timeout: 5 * time.Second,
	//Transport: netTransport,
}

type Caster interface {
	Cast(group *sync.WaitGroup, c chan core.ThingImp)
}
