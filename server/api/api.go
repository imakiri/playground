package api

import (
	"github.com/imakiri/playground/server/storage"
)

var Local storage.Local

func init() {
	Local = true
}
