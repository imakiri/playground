package utils

import (
	"github.com/imakiri/gorum/erres"
)

func IsNilSafe(l ...interface{}) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			return false
		}
	}
	return true
}

func IsNilSafeEx(l ...interface{}) (b []bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			b = append(b, false)
		} else {
			b = append(b, true)
		}
	}
	return
}

type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

type Service interface {
	Name() ServiceName
}

// FunctionID

type FunctionID uint16

func (e FunctionID) FID() uint16 {
	return uint16(e)
}

const FID_Detect FunctionID = 10
const FID_CreateUser FunctionID = 11

const FID_AuthLogin FunctionID = 0
const FID_AuthLogout FunctionID = 1

//

type ActionID uint64
type Meta struct {
	Error erres.Error
}
