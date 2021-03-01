package internal

import "github.com/imakiri/playground/core"

type ID struct {
	uuid  uint64
	pemid uint64
}

func (e ID) UUID() uint64 {
	return e.uuid
}

func (e ID) PemID() uint64 {
	return e.pemid
}

type Factor_Empty core.Factor

//

type Key_Login string

type Key_Password string

type Key_Random string

type Key_Email string

type Key_TelNum string

//

type Factor_Hash []byte

type Factor_Code string

//
