package data

import (
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/core"
	"github.com/jmoiron/sqlx"
)

func errorWrapper(err error) error {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return Error(core.CDataAlreadyExistError)
		case 1048:
			return Error(core.CAppDetecterIncorrectImageError)
		default:
			return Error(core.CDataInternalServiceError)
		}
	case error:
		break
	default:
		return Error(core.CUnknownError)
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return Error(core.CDataNotFoundError)
	default:
		return Error(core.CDataInternalServiceError)
	}
}

const query_CreateUser = "INSERT INTO main.users (login, name, avatar, pass_hash) VALUES (:login, :name, :avatar, :pass_hash)"
const query_GetUser = "SELECT name, avatar FROM main.users WHERE login = :login"
const query_GetPassHash = "SELECT pass_hash FROM main.users WHERE login = :login"
const query_UpdateUser = "UPDATE main.users SET name = :name, avatar = :avatar WHERE login = :login"
const query_UpdateUserPassHash = "UPDATE main.users SET pass_hash = :pass_hash WHERE login = :login"
const query_DeleteUser = "DELETE FROM main.users WHERE login = :login"

func (e *DB) CreateUser(c *core.ContainerCreateUser) error {
	q, args, err := sqlx.Named(query_CreateUser, c.Request)
	if err != nil {
		return errorWrapper(err)
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return errorWrapper(err)
	}

	return nil
}
func (e *DB) GetUser(c *core.ContainerGetUser) error {
	q, args, err := sqlx.Named(query_GetUser, c.Request)
	if err != nil {
		return errorWrapper(err)
	}

	q = e.db.Rebind(q)
	err = e.db.Get(&c.Response, q, args...)
	if err != nil {
		return errorWrapper(err)
	}

	return nil
}
func (e *DB) GetPassHash(c *core.ContainerGetUserPassHash) error {
	q, args, err := sqlx.Named(query_GetPassHash, c.Request)
	if err != nil {
		return errorWrapper(err)
	}

	q = e.db.Rebind(q)
	err = e.db.Get(&c.Response, q, args...)
	if err != nil {
		return errorWrapper(err)
	}

	return nil
}
func (e *DB) UpdateUser(c *core.ContainerUpdateUser) error {
	q, args, err := sqlx.Named(query_UpdateUser, c.Request)
	if err != nil {
		return errorWrapper(err)
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return errorWrapper(err)
	}

	return nil
}
func (e *DB) UpdateUserPassHash(c *core.ContainerUpdateUserPassHash) error {
	q, args, err := sqlx.Named(query_UpdateUserPassHash, c.Request)
	if err != nil {
		return errorWrapper(err)
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return errorWrapper(err)
	}

	return nil
}
func (e *DB) DeleteUser(c *core.ContainerDeleteUser) error {
	q, args, err := sqlx.Named(query_DeleteUser, c.Request)
	if err != nil {
		return errorWrapper(err)
	}

	_, err = e.db.Exec(q, args...)
	if err != nil {
		return errorWrapper(err)
	}

	return nil
}
