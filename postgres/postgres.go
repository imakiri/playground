package postgres

import (
	"database/sql"
	"github.com/imakiri/erres"
	"github.com/jmoiron/sqlx"
	"strings"
)

func NewDB(dsn string) (*sqlx.DB, error) {
	var db, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, erres.ConnectionError.Extend(0)
	}

	return db, err
}

// Wrapper for raw sql/sqlx/pgx error strings
func errWrapper(err error) *erres.Error {
	switch {
	case err == nil:
		return nil
	case err == sql.ErrTxDone:
		return erres.InternalServiceError.Extend(1).SetDescription(err.Error())
	}

	var e = err.Error()

	switch {

	case strings.Contains(e, "sqlx.bindNamedMapper: unsupported map type:"):
		return erres.InternalServiceError.Extend(1).SetDescription(err.Error())
	default:
		return erres.JustError.Extend(1).SetDescription(err.Error())
	}
}
