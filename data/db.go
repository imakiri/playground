package data

import (
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/core"
	"github.com/jmoiron/sqlx"
)

func errTypeCast(err error) (type_ string) {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return core.CDataAlreadyExistError
		case 1048:
			return core.CAppDetecterIncorrectImageError
		default:
			return core.CDataInternalServiceError
		}
	case error:
		break
	default:
		return core.CUnknownError
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return core.CDataNotFoundError
	default:
		return core.CDataInternalServiceError
	}
}

type DBMainUserId struct {
	Id uint `db:"id" json:"dataInternalMainUserId"`
}
type DBMainUserLogin struct {
	Login string `db:"login" json:"dataInternalMainUserLogin"`
}
type DBMainUserAvatar struct {
	Avatar []byte `db:"avatar" json:"dataInternalMainUserAvatar"`
}
type DBMainUserName struct {
	Name string `db:"name" json:"dataInternalMainUserName"`
}
type DBMainUserPassHash struct {
	PassHash []byte `db:"pass_hash" json:"dataInternalMainUserPassHash"`
}

const query_dataInternalMainCreateUser = "INSERT INTO main.users (login, name, avatar, pass_hash) VALUES (:login, :name, :avatar, :pass_hash)"
const query_dataInternalMainGetUser = "SELECT name, avatar FROM main.users WHERE login = :login"
const query_dataInternalMainGetPassHash = "SELECT pass_hash FROM main.users WHERE login = :login"
const query_dataInternalMainUpdateUser = "UPDATE main.users SET name = :name, avatar = :avatar WHERE login = :login"
const query_dataInternalMainUpdateUserPassHash = "UPDATE main.users SET pass_hash = :pass_hash WHERE login = :login"
const query_dataInternalMainDeleteUser = "DELETE FROM main.users WHERE login = :login"

type DBMainCreateUser struct {
	Request struct {
		DBMainUserLogin
		DBMainUserAvatar
		DBMainUserName
		DBMainUserPassHash
	}
}
type DBMainGetUser struct {
	Request struct {
		DBMainUserId
		DBMainUserLogin
	}
	Response struct {
		DBMainUserAvatar
		DBMainUserName
	}
}
type DBMainGetUserPassHash struct {
	Request struct {
		DBMainUserLogin
	}
	Response struct {
		DBMainUserPassHash
	}
}
type DBMainUpdateUser struct {
	Request struct {
		DBMainUserId
		DBMainUserLogin
		DBMainUserAvatar
		DBMainUserName
	}
}
type DBMainUpdateUserPassHash struct {
	Request struct {
		DBMainUserLogin
		DBMainUserPassHash
	}
}
type DBMainDeleteUser struct {
	Request struct {
		DBMainUserId
		DBMainUserLogin
	}
}

func (e *DB) DBMainCreateUser(c *DBMainCreateUser) error {
	q, args, err := sqlx.Named(query_dataInternalMainCreateUser, c.Request)
	if err != nil {
		return err
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (e *DB) DBMainGetUser(c *DBMainGetUser) error {
	q, args, err := sqlx.Named(query_dataInternalMainGetUser, c.Request)
	if err != nil {
		return err
	}

	q = e.db.Rebind(q)
	err = e.db.Get(&c.Response, q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (e *DB) DBMainGetPassHash(c *DBMainGetUserPassHash) error {
	q, args, err := sqlx.Named(query_dataInternalMainGetPassHash, c.Request)
	if err != nil {
		return err
	}

	q = e.db.Rebind(q)
	err = e.db.Get(&c.Response, q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (e *DB) DBMainUpdateUser(c *DBMainUpdateUser) error {
	q, args, err := sqlx.Named(query_dataInternalMainUpdateUser, c.Request)
	if err != nil {
		return err
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (e *DB) DBMainUpdateUserPassHash(c *DBMainUpdateUserPassHash) error {
	q, args, err := sqlx.Named(query_dataInternalMainUpdateUserPassHash, c.Request)
	if err != nil {
		return err
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (e *DB) DBMainDeleteUser(c *DBMainDeleteUser) error {
	q, args, err := sqlx.Named(query_dataInternalMainDeleteUser, c.Request)
	if err != nil {
		return err
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return err
	}

	return nil
}
