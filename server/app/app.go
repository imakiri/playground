package app

type Entity interface {
	GetSlice(q map[string]interface{})
	User
	Location
	Visit
}

type User interface {
}

type Location interface {
}

type Visit interface {
}
