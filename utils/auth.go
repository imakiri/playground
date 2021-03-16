package utils

type Salt string

type CredentialPlain string

type CredentialObscure []byte

type UUID uint64

type PEMID uint64

type ID struct {
	uuid  UUID
	pemid []PEMID
}

func (e ID) UUID() UUID {
	return e.uuid
}

func (e ID) PemID() []PEMID {
	return e.pemid
}

type Action interface{}
