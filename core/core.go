package core

type Config struct {
	Database struct {
		User     string
		Password string
		Address  string
		Port     string
	}
	DSN       string
	ApiKey    string
	Salt      string
	IPSDomain string
}

type Services struct {
	FaceDetection FaceDetecterClient
}

type Settings struct {
	Config
	Services
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
