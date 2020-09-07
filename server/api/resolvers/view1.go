package resolvers

import (
	"github.com/imakiri/playground/server/core"
	"io"
)

type v10 bool

func (v10) Resolve(parcel core.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View1 v0")
}

var V1 v10
