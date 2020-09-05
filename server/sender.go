package server

import "github.com/imakiri/playground/server/interfaces"

var Sender sender

type sender bool

func (sender) Send(api interfaces.Api, k string, c chan interfaces.Thing) {

}
