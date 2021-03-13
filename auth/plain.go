package auth

import "github.com/imakiri/gorum/core"

type Plain interface {
	Create(pemid core.PEMID, credential core.Credential) (core.UUID, error)
	Read(credential core.Credential) (core.UUID, core.PEMID, error)
	ReadCred(uuid core.UUID) (core.PEMID, core.Credential, error)
	Update(uuid core.UUID, pemid *core.PEMID, credential *core.Credential) error
	Delete(uuid core.UUID) error
}
