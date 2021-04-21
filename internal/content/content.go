package content

import (
	"github.com/imakiri/gorum/internal/types"
)

type ServiceAvatar interface {
	Get128(userUUID types.ModelUserUUID, container types.ModelUserAvatar128) error
	Get256(userUUID types.ModelUserUUID, container types.ModelUserAvatar256) error
	Get512(userUUID types.ModelUserUUID, container types.ModelUserAvatar512) error
	Set(update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error
}
