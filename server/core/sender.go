package core

var SenderImp sender

type sender bool

func (sender) Send(api Api, k string, c chan Thing) {

}
