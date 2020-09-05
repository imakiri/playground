package server

type Thing struct {
	Header string
	Data   []byte
	Error  error
}

func (t *Thing) GetHeader() string {
	return t.Header
}

func (t *Thing) GetData() []byte {
	return t.Data
}

func (t *Thing) GetError() error {
	return t.Error
}
