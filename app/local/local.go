package local

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User interface {
	Add()
	Auth()
	Delete()
}

type View interface {
	AllUsers()
	AllPicturesByUser()
}

type SearchHistory interface {
}

type AddUserPicture interface {
	AddUserPicture()
}

type DeleteUserPicture interface {
	DeleteUserPicture()
}
