package types

type (
	Page struct {
		Start  uint64
		Length uint64
	}
)

type (
	ModelUserUUID      string
	ModelUserPemID     int16
	ModelUserNickName  string
	ModelUserFullName  string
	ModelUserPosts     int32
	ModelUserAvatar512 []byte
	ModelUserAvatar256 []byte
	ModelUserAvatar128 []byte

	ModelDate    int64
	ModelContent string

	ModelPostUUID string

	ModelThreadUUID string
	ModelThreadName string

	ModelCategoryUUID string
	ModelCategoryName string

	ModelCookieKey string

	ModelLogpassLogin    []byte
	ModelLogpassPassword []byte
)

type (
	ModelUser struct {
		UserUUID         ModelUserUUID
		RegistrationDate ModelDate
		Nickname         ModelUserNickName
		Fullname         ModelUserFullName
		Avatar512        ModelUserAvatar512
		Avatar256        ModelUserAvatar256
		Avatar128        ModelUserAvatar128
	}
	ModelThread struct {
		ThreadUUID   ModelThreadUUID
		CategoryUUID ModelCategoryUUID
		UserUUID     ModelUserUUID
		Name         ModelThreadName
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Header       ModelContent
	}
	ModelPost struct {
		PostUUID     ModelPostUUID
		ThreadUUID   ModelThreadUUID
		UserUUID     ModelUserUUID
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Content      ModelContent
	}
	ModelCategory struct {
		CategoryUUID ModelCategoryUUID
		Name         ModelCategoryName
	}
	ModelCookie struct {
		Key            ModelCookieKey
		UserUUID       ModelUserUUID
		PemID          ModelUserPemID
		ExpirationDate ModelDate
	}
	ModelLogpass struct {
		UserUUID ModelUserUUID
		Login    ModelLogpassLogin
		Password ModelLogpassPassword
		PemID    ModelUserPemID
	}
)
