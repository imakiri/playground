package data

import (
	"github.com/imakiri/playground/core"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

var ConnectionInternalMain core.DataInternalMain

func init() {
	var err error

	ConnectionInternalMain.SQLX_DB, err = sqlx.Connect("pgx", core.Conf.DSN)
	if err != nil {
		panic(err)
	}

	err = ConnectionInternalMain.SQLX_DB.Ping()
	if err != nil {
		panic(err)
	}
}
