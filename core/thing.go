package core

type ThingImp struct {
	Header string
	Data   []byte
	Error  error
}

func (t *ThingImp) GetHeader() string {
	return t.Header
}

func (t *ThingImp) GetData() []byte {
	return t.Data
}

func (t *ThingImp) GetError() error {
	return t.Error
}
