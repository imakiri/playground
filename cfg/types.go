package cfg

type Config struct {
	Databases ConfigDatabases
}

type ConfigDatabases struct {
	Mongo    ConfigDatabaseMongo
	Postgres ConfigDatabasePostgres
}

type ConfigDatabaseMongo struct {
}

type ConfigDatabasePostgres struct {
}

type Secret struct {
	Databases SecretDatabases
}

type SecretDatabases struct {
	Mongo    SecretDatabaseMongo
	Postgres SecretDatabasePostgres
}

type SecretDatabasePostgres struct {
	DSN string
}

type SecretDatabaseMongo struct {
	DSN string
}
