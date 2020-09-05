package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"io"
)

type v20 bool

func (v20) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View2 v0")
}

var V2 v20
