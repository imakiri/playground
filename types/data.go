package types

type (
	ModelUserUUID             string
	ModelUserPemID            int16
	ModelUserNickName         string
	ModelUserFullName         string
	ModelUserPosts            int32
	ModelUserAvatar512        []byte
	ModelUserAvatar256        []byte
	ModelUserAvatar128        []byte
	ModelUserRegistrationDate int64

	ModelDate    int64
	ModelContent string

	ModelPostUUID string

	ModelThreadUUID string
	ModelThreadName string

	ModelCategoryUUID string
	ModelCategoryName string

	ModelCookieKey            string
	ModelCookieExpirationDate int64

	ModelLogpassLogin    []byte
	ModelLogpassPassword []byte
)

type (
	ModelUser struct {
		UserUUID         ModelUserUUID
		RegistrationDate ModelUserRegistrationDate
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
		UUID           ModelUserUUID
		PemID          ModelUserPemID
		ExpirationDate ModelCookieExpirationDate
	}
	ModelLogpass struct {
		UUID     ModelUserUUID
		Login    ModelLogpassLogin
		Password ModelLogpassPassword
		PemID    ModelUserPemID
	}
)

type (
	ViewUserID struct {
		UserUUID ModelUserUUID
		PemID    ModelUserPemID
	}
	ViewUserAvatar struct {
		Avatar512 ModelUserAvatar512
		Avatar256 ModelUserAvatar256
		Avatar128 ModelUserAvatar128
	}
)

type (
	ViewCookieByUUID struct {
		PemID          ModelUserPemID
		Key            ModelCookieKey
		ExpirationDate ModelCookieExpirationDate
	}
	ViewLogpassByUUID struct {
		PemID    ModelUserPemID
		Login    ModelLogpassLogin
		Password ModelLogpassPassword
	}
	ViewLogpassByLogin struct {
		UUID     ModelUserUUID
		PemID    ModelUserPemID
		Password ModelLogpassPassword
	}
)

type (
	ViewUserProfile struct {
		RegistrationDate ModelUserRegistrationDate
		NickName         ModelUserNickName
		FullName         ModelUserFullName
		Avatar512        ModelUserAvatar512
	}
	ViewUserProfileFromThread struct {
		UserUUID         ModelUserUUID
		RegistrationDate ModelUserRegistrationDate
		NickName         ModelUserNickName
		FullName         ModelUserFullName
		Avatar256        ModelUserAvatar256
	}
	ViewUserProfileFromMain struct {
		UserUUID  ModelUserUUID
		NickName  ModelUserNickName
		Avatar128 ModelUserAvatar128
	}
	ViewUserProfileUpdate struct {
		NickName *ModelUserNickName
		FullName *ModelUserFullName
		Avatar   *ViewUserAvatar
	}
	ViewPostCreate struct {
		UserUUID   ModelUserUUID
		ThreadUUID ModelThreadUUID
		Content    ModelContent
	}
	ViewPostUpdate struct {
		Content ModelContent
	}
	ViewPostByThreadUUID struct {
		PostUUID     ModelPostUUID
		UserUUID     ModelUserUUID
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Content      ModelContent
	}
	ViewThreadCreate struct {
		CategoryUUID ModelCategoryUUID
		UserUUID     ModelUserUUID
		Name         ModelThreadName
		Header       ModelContent
	}
	ViewThread struct {
		Category     ModelCategory
		Author       ViewUserProfileFromThread
		Name         ModelThreadName
		DateAdded    ModelDate
		DateLastEdit ModelDate
		Header       ModelContent
		Content      ViewThreadContent
	}
	ViewThreadContent struct {
		Users []ViewUserProfileFromThread
		Posts []ViewPostByThreadUUID
	}
	ViewThreadsByCategory struct {
		ThreadUUID ModelThreadUUID
		Name       ModelThreadName
	}
	ViewThreadUpdate struct {
		CategoryUUID *ModelCategoryUUID
		Name         *ModelThreadName
		Header       *ModelContent
	}
)
