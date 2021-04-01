package types

type Secret struct {
	Postgres SecretPostgres
}

type SecretPostgres struct {
	DSN string
}
