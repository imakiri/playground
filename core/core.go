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

type Identity struct{}

type Checker interface {
	For(funcName FuncName) (bool, error)
}

type FuncName uint8
type UserGroup string
type Permission uint
type Permissions map[FuncName]bool
type PermissionsTable map[UserGroup]Permissions
type PersonalPermissionsTable map[FieldUserId]Permissions

const FN_Detect FuncName = 0
const FN_CreateUser FuncName = 1

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
const CAccessDenied = "AccessDenied"
const CSerializationError = "SerializationError"
const CUnknownError = "UnknownError"

type FieldUserId struct {
	Id uint `db:"id" json:"dataInternalMainUserId"`
}
type FieldUserLogin struct {
	Login string `db:"login" json:"dataInternalMainUserLogin"`
}
type FieldUserAvatar struct {
	Avatar []byte `db:"avatar" json:"dataInternalMainUserAvatar"`
}
type FieldUserName struct {
	Name string `db:"name" json:"dataInternalMainUserName"`
}
type FieldUserPassHash struct {
	PassHash []byte `db:"pass_hash" json:"dataInternalMainUserPassHash"`
}

type ContainerCreateUser struct {
	Request struct {
		FieldUserLogin
		FieldUserAvatar
		FieldUserName
		FieldUserPassHash
	}
}
type ContainerGetUser struct {
	Request struct {
		FieldUserId
		FieldUserLogin
	}
	Response struct {
		FieldUserAvatar
		FieldUserName
	}
}
type ContainerGetUserPassHash struct {
	Request struct {
		FieldUserLogin
	}
	Response struct {
		FieldUserPassHash
	}
}
type ContainerUpdateUser struct {
	Request struct {
		FieldUserId
		FieldUserLogin
		FieldUserAvatar
		FieldUserName
	}
}
type ContainerUpdateUserPassHash struct {
	Request struct {
		FieldUserLogin
		FieldUserPassHash
	}
}
type ContainerDeleteUser struct {
	Request struct {
		FieldUserId
		FieldUserLogin
	}
}
