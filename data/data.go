package data

// Model's types ---------------------------------------------------------------------------------------------------------
type (

	// Model.User types ------------------------------------------------------------------------------------------------

	ModelUserUUID             string
	ModelUserPemID            int16
	ModelUserNickname         string
	ModelUserFullname         string
	ModelUserAvatar           []byte
	ModelUserRegistrationDate int64

	// Model.Post types ------------------------------------------------------------------------------------------------

	ModelPostUUID     uint64
	ModelPostDate     int64
	ModelPostLastEdit int64
	ModelPostContent  string

	// Model.Thread types ----------------------------------------------------------------------------------------------

	ModelThreadUUID         string
	ModelThreadCreationDate int64
	ModelThreadTopic        int16
	ModelThreadName         string

	// Model.Cookie types ----------------------------------------------------------------------------------------------

	ModelCookieKey            string
	ModelCookieExpirationDate int64

	// Model.Logpass types ---------------------------------------------------------------------------------------------

	ModelLogpassLogin    []byte
	ModelLogpassPassword []byte
)

// Model types ---------------------------------------------------------------------------------------------------------
type (
	ModelUser struct {
		UUID             ModelUserUUID
		Nickname         ModelUserNickname
		Fullname         ModelUserFullname
		Avatar           ModelUserAvatar
		RegistrationDate ModelUserRegistrationDate
	}
	ModelThread struct {
		UUID         ModelThreadUUID
		UserUUID     ModelUserUUID
		CreationDate ModelThreadCreationDate
		Topic        ModelThreadTopic
		Name         ModelThreadName
	}
	ModelPost struct {
		ThreadUUID ModelThreadUUID
		UserUUID   ModelUserUUID
		Date       ModelPostDate
		LastEdit   ModelPostLastEdit
		Content    ModelPostContent
		UUID       ModelPostUUID
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

// Custom types --------------------------------------------------------------------------------------------------------
type (
	UserID struct {
		UUID  ModelUserUUID
		PemID ModelUserPemID
	}
)
