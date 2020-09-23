package db

func (re *Re) Read(p []byte) (n int, err error) {

	return
}

func (re *Re) Write(p []byte) (n int, err error) {

	return
}

func (re *Re) Error() string {
	return re.Err.Error()
}

func (re Re) Check(err error, c chan Re) bool {
	if err != nil {
		re.Err = err
		c <- re
		return true
	}
	return false
}
