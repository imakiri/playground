package core

type ConfigData struct {
	DSN string
}
type ConfigApp struct {
	HashCost int
	Salt     string
}
type ConfigAPI struct {
	ToStart bool
}
type ConfigUI struct {
	ToStart   bool
	IPSDomain string
}
type Config struct {
	Data ConfigData
	App  ConfigApp
	API  ConfigAPI
	UI   ConfigUI
}

type Settings struct {
	Config
}
