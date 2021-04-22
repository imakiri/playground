package postgres

import (
	"github.com/imakiri/gorum/internal/cfg"
	"testing"
)

func TestNewConnection(t *testing.T) {
	var config = cfg.ConfigDatabasePostgres{
		Host:    "imakiri-ips.ddns.net",
		Port:    5432,
		Dbname:  "main",
		Sslmode: "disable",
	}
	var secret = cfg.SecretDatabasePostgres{
		User:     "test",
		Password: "testpass",
	}
	var _, err = NewConnection(nil, config, secret)
	if err != nil {
		t.Error(err)
	}
}
