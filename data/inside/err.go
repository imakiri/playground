package inside

type InternalServiceError struct {
	err string
}

func (e InternalServiceError) Error() string {
	return e.err
}

type IncorrectArgumentError struct {
	err string
}

func (e IncorrectArgumentError) Error() string {
	return e.err
}

type NotFoundError struct {
	err string
}

func (e NotFoundError) Error() string {
	return e.err
}
