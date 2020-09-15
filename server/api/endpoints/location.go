package endpoints

import (
	"github.com/imakiri/playground/server/core"
	"io"
)

type l0 bool

func (l0) Resolve(parcel core.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View2 v0")
}

var Location l0
