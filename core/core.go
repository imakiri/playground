package core

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf Config

func init() {
	var err error
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Database struct {
		User     string
		Password string
		Address  string
		Port     string
	}
	DSN    string
	ApiKey string
	Salt   string
}

type Header struct {
	ActionID uint `json:"action_id"`
}

type Package struct {
	Session `json:"session"`
	Header  `json:"header"`
	Status  `json:"status"`
}

type Status struct {
	Type_ string `json:"type"`
	Value string `json:"value"`
}

func (e Status) Error() string {
	return fmt.Sprintf("%s: [%s]", e.Type_, e.Value)
}
func (e Status) Type() string {
	return e.Type_
}
func (e Status) IsOK() bool {
	if e.Type_ == CStatusOK {
		return true
	} else {
		return false
	}
}

type StatusPointer interface {
	StatusPtr() *Status
}

type StatusInt interface {
	error
	Type() string
}

func StatusType2StrCaster(s StatusInt) string {
	switch s.(type) {
	case StatusOk:
		return CStatusOK
	case SerializationError:
		return CSerializationError
	case DataAlreadyExistError:
		return CDataAlreadyExistError
	case DataIncorrectArgumentError:
		return CDataIncorrectArgumentError
	case DataInternalServiceError:
		return CDataInternalServiceError
	case DataNotFoundError:
		return CDataNotFoundError
	case AppDetecterIncorrectImageError:
		return CAppDetecterIncorrectImageError
	case WebIcoInitError:
		return CWebIcoInitError
	case WebCssReadError:
		return CWebCssReadError
	case WebTemplateParseError:
		return CWebTemplateParseError
	case WebTemplateExecuteError:
		return CWebTemplateExecuteError
	default:
		return CUnknownError
	}
}

func NewStatus(e StatusInt, err error) StatusInt {
	if err == nil {
		return StatusOk{
			Status{
				Type_: CStatusOK,
				Value: "",
			},
		}
	} else {
		switch e.(type) {
		case DataAlreadyExistError:
			return DataAlreadyExistError{Status{
				Type_: CDataAlreadyExistError,
				Value: err.Error(),
			}}
		case DataIncorrectArgumentError:
			return DataIncorrectArgumentError{Status{
				Type_: CDataIncorrectArgumentError,
				Value: err.Error(),
			}}
		case DataInternalServiceError:
			return DataInternalServiceError{Status{
				Type_: CDataInternalServiceError,
				Value: err.Error(),
			}}
		case DataNotFoundError:
			return DataNotFoundError{Status{
				Type_: CDataNotFoundError,
				Value: err.Error(),
			}}
		case AppDetecterIncorrectImageError:
			return AppDetecterIncorrectImageError{Status{
				Type_: CAppDetecterIncorrectImageError,
				Value: err.Error(),
			}}
		case WebIcoInitError:
			return WebIcoInitError{Status{
				Type_: CWebIcoInitError,
				Value: err.Error(),
			}}
		case WebCssReadError:
			return WebCssReadError{Status{
				Type_: CWebCssReadError,
				Value: err.Error(),
			}}
		case WebTemplateParseError:
			return WebTemplateParseError{Status{
				Type_: CWebTemplateParseError,
				Value: err.Error(),
			}}
		case WebTemplateExecuteError:
			return WebTemplateExecuteError{Status{
				Type_: CWebTemplateExecuteError,
				Value: err.Error(),
			}}
		default:
			return UnknownError{Status{
				Type_: CUnknownError,
				Value: err.Error(),
			}}
		}
	}
}

type StatusOk struct{ Status }
type SerializationError struct{ Status }
type UnknownError struct{ Status }

func (e StatusOk) Type() string {
	return CStatusOK
}
func (e SerializationError) Type() string {
	return CSerializationError
}
func (e UnknownError) Type() string {
	return CUnknownError
}

const CDataInternalServiceError = "DataInternalServiceError"
const CDataIncorrectArgumentError = "DataIncorrectArgumentError"
const CDataNotFoundError = "DataNotFoundError"
const CDataAlreadyExistError = "DataAlreadyExistError"

const CAppDetecterIncorrectImageError = "AppDetecterIncorrectImageError"

const CWebIcoInitError = "WebIcoInitError"
const CWebCssReadError = "WebCssReadError"
const CWebTemplateParseError = "WebTemplateParseError"
const CWebTemplateExecuteError = "WebTemplateExecuteError"

const CStatusOK = "StatusOK"
const CSerializationError = "SerializationError"
const CUnknownError = "UnknownError"
