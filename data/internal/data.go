package internal

import "github.com/imakiri/playground/data/schema"

func (r) AddData(d schema.Data) (err error) {
	return
}

func (r) DeleteData(id uint) (err error) {
	return
}

func (r) SearchForData(q string) (re []schema.Data) {
	return
}

func (r) GetLastUserData(login string, size uint8) (re []schema.Data) {
	return
}

func (r) GetLastData(size uint8) (re []schema.Data) {
	return
}
