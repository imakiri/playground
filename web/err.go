package web

type ERROR string

func (e ERROR) Error() string {
	return string(e)
}

type ERROR_InitIco struct {
	ERROR
}
type ERROR_ReadCss struct {
	ERROR
}
type ERROR_ParseTemplate struct {
	ERROR
}
type ERROR_ExecuteTemplate struct {
	ERROR
}
