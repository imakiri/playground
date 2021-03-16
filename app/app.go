package app

import (
	"github.com/imakiri/gorum/data"
	"github.com/imakiri/gorum/transport"
	"github.com/imakiri/gorum/utils"
	"github.com/jackc/pgx/v4"
)

//type User struct {
//	db       *pgx.Conn
//	log      utils.LogService
//	config   *transport.App
//	configDB *transport.Data
//}
//
//func NewService(c *transport.System) (*User, error) {
//	var s User
//	var err error
//
//	s.config = c.GetApp()
//	s.configDB = c.GetData()
//
//	s.db, err = data.Connect(c.GetData())
//	if err != nil {
//		return nil, err
//	}
//
//	return &s, err
//}
