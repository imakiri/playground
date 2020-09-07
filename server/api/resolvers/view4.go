package resolvers

import (
	"github.com/imakiri/playground/server/core"
	"io"
)

type v40 bool

func (v40) Resolve(parcel core.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View4 v0")
}

var V4 v40
