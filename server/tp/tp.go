package tp

type Api interface {
	GetSlice(p Parcel) *Parcel
}

type Get interface {
	Get(string) string
}

type Parcel struct {
	header map[string]string
	body   interface{}
	error  uint8
}

type Status struct {
	key   uint8
	value string
}

func (s *Status) Error() string {
	return string(s.key) + s.value
}

type create bool

var Create create = true

func (c create) Parcel(header map[string]string, body interface{}) (p Parcel) {
	if c == false {
		p = Parcel{error: 1}
		return
	}
	p = Parcel{header: header, body: body, error: 0}

	return
}

//func (c create) Aberration(key uint8, value string) Parcel {
//	return Status{key, value}
//}

//type Create interface {
//	Parcel(header map[uint8]string, body interface{}) Parcel
//	Status(key uint8, value string) Status
//}

func init() {}
