package internal

import "github.com/imakiri/playground/data/schema"

func (consecutive) AddData(d schema.Data) (err error) {
	return
}

func (consecutive) DeleteData(id uint) (err error) {
	return
}

func (consecutive) SearchForData(q string) (re []schema.Data) {
	return
}

func (consecutive) GetLastUserData(login string, size uint8) (re []schema.Data) {
	return
}

func (consecutive) GetLastData(size uint8) (re []schema.Data) {
	return
}
