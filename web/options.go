package web

import (
	"github.com/imakiri/gorum/web/internal"
)

type setRegistrar string

func (setRegistrar) Main() registrar {
	return internal.RegistrarMain
}

const Registrars setRegistrar = "the set of available registrars"

type setRedirector string

const Redirectors setRedirector = "the set of available redirectors"
