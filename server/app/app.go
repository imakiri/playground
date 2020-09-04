package app

import "github.com/imakiri/playground/server/api"

type App interface {
	MatchUp(parcel api.Parcel)
}
