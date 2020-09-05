package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"io"
)

type v00 bool

func (v00) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View0 v0")
}

type v01 bool

func (v01) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done View0 v1")
}

var V0 v01
