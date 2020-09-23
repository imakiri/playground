package endpoints

import "github.com/imakiri/playground/core"

type Resolver interface {
	Resolve(p core.Parcel)
}
