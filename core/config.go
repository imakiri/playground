package core

type ConfigDB struct {
	DSN string
}
type ConfigApp struct {
	HashCost int
	Salt     string
}
type ConfigAuth struct {
}
type ConfigContent struct {
}
type ConfigGate struct {
}
type ConfigAPI struct {
	ToStart bool
}
type ConfigUI struct {
	ToStart   bool
	IPSDomain string
}
type Config struct {
	DB      ConfigDB
	App     ConfigApp
	Auth    ConfigAuth
	Content ConfigContent
	Gate    ConfigGate
	API     ConfigAPI
	UI      ConfigUI
}

type Settings struct {
	Config Config
}
