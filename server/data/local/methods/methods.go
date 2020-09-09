package methods

type Methods struct{}

func (Methods) AddUser() {
	AddUser.AddUser()
}

func (Methods) AuthUser() {
	AuthUser.AuthUser()
}

func (Methods) DeleteUser() {
	DeleteUser.DeleteUser()
}

func (Methods) GetAllUsers() {
	GetAllUsers.GetAllUsers()
}

func (Methods) AddUserPicture() {
	AddUserPicture.AddUserPicture()
}

func (Methods) DeleteUserPicture() {
	DeleteUserPicture.DeleteUserPicture()
}

func (Methods) GetAllPicturesByUser() {
	GetAllPicturesByUser.GetAllPicturesByUser()
}
