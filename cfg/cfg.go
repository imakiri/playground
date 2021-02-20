package cfg

type DB struct {
	DSN string
}
type App struct {
}
type Auth struct {
	HashCost int
	Salt     string
}
type Content struct {
}
type System struct {
	DB      DB
	App     App
	Auth    Auth
	Content Content
}

type Gate struct {
}

type API struct {
}
type Web struct {
	LaunchRedir bool
	IPSDomain   string
}
type EI struct {
	API API
	Web Web
}

type Config struct {
	System System
	Gate   Gate
	EI     EI
}
