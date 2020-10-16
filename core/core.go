package core

import (
	"fmt"
	"time"
)

type Header struct {
	From string               `json:"from"`
	To   string               `json:"to"`
	Path map[string]time.Time `json:"path"`
}

type Package struct {
	Permissions `json:"permissions"`
	Header      `json:"header"`
	Err         `json:"error"`
}

type ERROR struct {
	Type_ string `json:"type"`
	Value string `json:"value"`
}

func (e ERROR) Error() string {
	return fmt.Sprintf("%s: [%s]", e.Type_, e.Value)
}
func (e ERROR) Type() string {
	return e.Type_
}

type Err interface {
	error
	Type() string
}

func NewError(e Err, value string) (err Err) {
	switch e.(type) {
	case DataAlreadyExistError:
		return DataAlreadyExistError{ERROR{
			Type_: CDataAlreadyExistError,
			Value: value,
		}}
	case DataIncorrectArgumentError:
		return DataIncorrectArgumentError{ERROR{
			Type_: CDataIncorrectArgumentError,
			Value: value,
		}}
	case DataInternalServiceError:
		return DataInternalServiceError{ERROR{
			Type_: CDataInternalServiceError,
			Value: value,
		}}
	case DataNotFoundError:
		return DataNotFoundError{ERROR{
			Type_: CDataNotFoundError,
			Value: value,
		}}
	case WebIcoInitError:
		return WebIcoInitError{ERROR{
			Type_: CWebIcoInitError,
			Value: value,
		}}
	case WebCssReadError:
		return WebCssReadError{ERROR{
			Type_: CWebCssReadError,
			Value: value,
		}}
	case WebTemplateParseError:
		return WebTemplateParseError{ERROR{
			Type_: CWebTemplateParseError,
			Value: value,
		}}
	case WebTemplateExecuteError:
		return WebTemplateExecuteError{ERROR{
			Type_: CWebTemplateExecuteError,
			Value: value,
		}}
	default:
		return UnknownError{ERROR{
			Type_: CUnknownError,
			Value: value,
		}}
	}
}

type UnknownError struct{ ERROR }

func (e UnknownError) Type() string {
	return CUnknownError
}

const CDataInternalServiceError = "DataInternalServiceError"
const CDataIncorrectArgumentError = "DataIncorrectArgumentError"
const CDataNotFoundError = "DataNotFoundError"
const CDataAlreadyExistError = "DataAlreadyExistError"

const CWebIcoInitError = "WebIcoInitError"
const CWebCssReadError = "WebCssReadError"
const CWebTemplateParseError = "WebTemplateParseError"
const CWebTemplateExecuteError = "WebTemplateExecuteError"

const CUnknownError = "UnknownError"
