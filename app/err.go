package app

type ERROR string

func (b ERROR) Error() string {
	return string(b)
}

type NotAuthorizedError struct{ ERROR }
type InternalServiceError struct{ ERROR }
type IncorrectImageError struct{ ERROR }
