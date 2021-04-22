package postgres

import (
	"crypto/tls"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/cfg"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type Connection struct {
	db *sqlx.DB
}

func NewConnection(log pgx.Logger, config cfg.ConfigDatabasePostgres, secret cfg.SecretDatabasePostgres) (*Connection, error) {
	var tlsConf *tls.Config
	if config.Sslmode == "enable" {
		//
	}

	var conf, err = pgx.ParseConfig("")
	if err != nil {
		return nil, err
	}

	conf.Config.Host = config.Host
	conf.Config.Port = config.Port
	conf.Config.Database = config.Dbname
	conf.Config.User = secret.User
	conf.Config.Password = secret.Password
	conf.Config.TLSConfig = tlsConf
	conf.Config.ConnectTimeout = 5 * time.Second
	conf.Logger = log

	var db = stdlib.OpenDB(*conf)
	if db == nil {
		return nil, erres.ConnectionError
	}

	var conn = new(Connection)
	conn.db = sqlx.NewDb(db, "pgx")

	if err = conn.db.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
