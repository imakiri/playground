package types

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
	ViewCookie struct {
		PemID          ModelUserPemID
		Key            ModelCookieKey
		ExpirationDate ModelDate
	}
	ViewLogpass struct {
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
		RegistrationDate ModelDate
		NickName         ModelUserNickName
		FullName         ModelUserFullName
		Avatar512        ModelUserAvatar512
	}
	ViewUserProfileFromThread struct {
		UserUUID         ModelUserUUID
		RegistrationDate ModelDate
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
