package resolvers

import "github.com/imakiri/playground/server/core"

type Resolver interface {
	Resolve(p core.Parcel)
}
