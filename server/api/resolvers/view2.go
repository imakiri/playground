package resolvers

import (
	"github.com/imakiri/playground/server/core"
	"io"
)

type v20 bool

func (v20) Resolve(parcel core.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View2 v0")
}

var V2 v20
