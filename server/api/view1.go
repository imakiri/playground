package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"io"
)

type v10 bool

func (v10) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View1 v0")
}

var V1 v10
