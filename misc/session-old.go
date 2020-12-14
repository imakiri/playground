package misc

type UserPassRaw struct {
	PassRaw string `json:"passRaw"`
}

type Session struct {
	id string
	Permissions
}

type Permissions struct {
}

type Authorization struct {
	Request struct {
		DataInternalMainUserLogin
		UserPassRaw
	}
	Response struct {
		Session
	}
}
