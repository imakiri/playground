package cfg

type TypeDB string

const (
	MONGO    TypeDB = "mongo"
	POSTGRES TypeDB = "postgres"
)

type Config struct {
	Connections ConfigConnections
	Services    ConfigServices
}

type ConfigConnections struct {
	Mongo    map[string]ConfigDatabaseMongo    `yaml:"mongo"`
	Postgres map[string]ConfigDatabasePostgres `yaml:"postgres"`
}

type ConfigDatabaseMongo struct {
}

type ConfigDatabasePostgres struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	Dbname  string `yaml:"dbname"`
	Sslmode string `yaml:"sslmode"`
}

type Secret struct {
	Connections SecretConnections
}

type SecretConnections struct {
	Mongo    map[string]SecretDatabaseMongo    `yaml:"mongo"`
	Postgres map[string]SecretDatabasePostgres `yaml:"postgres"`
}

type SecretDatabasePostgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type SecretDatabaseMongo struct {
}

type ConfigServices struct {
	Avatar ConfigService
}

type ConfigService struct {
	DBType   TypeDB `yaml:"dbtype"`
	ConnName string `yaml:"connname"`
}
