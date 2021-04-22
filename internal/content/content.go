package content

import (
	"github.com/imakiri/gorum/internal/auth"
	"github.com/imakiri/gorum/internal/postgres"
)

type ServiceAvatar interface {
	Get128(userID auth.UserID, container []byte) error
	Get256(userID auth.UserID, container []byte) error
	Get512(userID auth.UserID, container []byte) error
	Set(update bool, userID auth.UserID, avatar postgres.ViewUserAvatar) error
}
