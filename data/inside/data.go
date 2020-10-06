package inside

import "github.com/imakiri/playground/data/schema"

func (R) AddData(d schema.Data) (err error) {
	return
}

func (R) DeleteData(id uint) (err error) {
	return
}

func (R) SearchForData(q string) (re []schema.Data) {
	return
}

func (R) GetLastUserData(login string, size uint8) (re []schema.Data) {
	return
}

func (R) GetLastData(size uint8) (re []schema.Data) {
	return
}
