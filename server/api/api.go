package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"github.com/imakiri/playground/server/storage"
	"io"
)

var Local storage.Local

type v0 bool

func (v0) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do0")
}

var V0 v0

type v1 bool

func (v1) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do1")
}

var V1 v1

type v2 bool

func (v2) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do2")
}

var V2 v2

type v3 bool

func (v3) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do3")
}

var V3 v3

type v4 bool

func (v4) Resolve(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do4")
}

var V4 v4

func init() {
	Local = true
}
