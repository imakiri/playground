package resolvers

import (
	"github.com/imakiri/playground/server/core"
	"io"
)

type v30 bool

func (v30) Resolve(parcel core.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View3 v0")
}

var V3 v30
