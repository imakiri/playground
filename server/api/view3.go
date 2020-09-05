package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"io"
)

type v30 bool

func (v30) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View3 v0")
}

var V3 v30
