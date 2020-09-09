package local

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Methods interface {
	addUser
	AuthUser
	DeleteUser
	GetAllUsers
	AddUserPicture
	DeleteUserPicture
	GetAllPicturesByUser
}

type addUser interface {
	AddUser()
}

type AuthUser interface {
	AuthUser()
}

type DeleteUser interface {
	DeleteUser()
}

type GetAllUsers interface {
	GetAllUsers()
}

type AddUserPicture interface {
	AddUserPicture()
}

type DeleteUserPicture interface {
	DeleteUserPicture()
}

type GetAllPicturesByUser interface {
	GetAllPicturesByUser()
}
