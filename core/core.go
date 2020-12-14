package core

type Data struct {
	DSN    string
	ApiKey string
}
type App struct {
	HashCost int
	Salt     string
}
type API struct {
	ToStart bool
}
type Web struct {
	ToStart   bool
	IPSDomain string
}
type Config struct {
	Data
	App
	API
	Web
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
