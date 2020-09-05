package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"io"
)

type v40 bool

func (v40) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View4 v0")
}

var V4 v40
