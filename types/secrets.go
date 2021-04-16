package types

type Secret struct {
	Data    SecretData
	Content SecretContent
}

type SecretData struct {
	Postgres SecretDataPostgres
}
type SecretContent struct {
	Postgres SecretDataPostgres
}

type SecretDataPostgres struct {
	DSN string
}
