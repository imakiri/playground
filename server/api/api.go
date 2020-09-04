package api

import (
	"github.com/imakiri/playground/server/interfaces"
	"github.com/imakiri/playground/server/storage"
	"io"
)

var Local storage.Local

type App bool

func (App) Do0(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do0")
}

func (App) Do1(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do1")
}

func (App) Do2(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do2")
}

func (App) Do3(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do3")
}

func (App) Do4(parcel interfaces.Parcel) {
	_, _ = io.WriteString(parcel.ResponseWriter, "Done Do4")
}

func init() {
	Local = true
}
