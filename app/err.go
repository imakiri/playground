package app

type ERROR string

func (b ERROR) Error() string {
	return string(b)
}

type ERROR_NotAuthorized struct{ ERROR }
