package core

type UserPassRaw struct {
	PassRaw string `json:"passRaw"`
}

type Permissions struct {
}

type Authorization struct {
	Request struct {
		DataInternalMainUserLogin
		UserPassRaw
	}
	Response struct {
		Permissions
	}
}
