package auth

import "github.com/imakiri/gorum/utils"

type Plain interface {
	Create(pemid utils.PEMID, credential utils.Credential) (utils.UUID, error)
	Read(credential utils.Credential) (utils.UUID, utils.PEMID, error)
	ReadCred(uuid utils.UUID) (utils.PEMID, utils.Credential, error)
	Update(uuid utils.UUID, pemid *utils.PEMID, credential *utils.Credential) error
	Delete(uuid utils.UUID) error
}
