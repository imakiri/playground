package data

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/misc"
)

func errorWrapper(err error) error {
	switch err.(type) {
	case *mysql.MySQLError:
		e := err.(*mysql.MySQLError)
		switch e.Number {
		case 1062:
			return core.Status_AlreadyExist
		case 1048:
			return core.Status_UnknownError
		default:
			return core.Status_InternalServiceError
		}
	case error:
		break
	default:
		return core.Status_UnknownError
	}

	switch st := err.Error(); st {
	case "sql: no rows in result set":
		return core.Status_NotFound
	default:
		return core.Status_InternalServiceError
	}
}

const query_CreateUser = "INSERT INTO main.user_credentials (login, pass_hash, group_id) VALUES (:login, :pass_hash, :group_id)"
const query_GetUserInfo = "SELECT name, avatar FROM main.users WHERE login = :login"
const query_UpdateUserInfo = "UPDATE main.users SET name = :name, avatar = :avatar WHERE login = :login"
const query_GetUserList = ""

const query_GetThreadContent = ""
const query_GetPostInfo = ""
const query_PostToThread = ""

const query_GetPassHash = "SELECT pass_hash FROM main.users WHERE login = :login"
const query_UpdateUserPassHash = "UPDATE main.users SET pass_hash = :pass_hash WHERE login = :login"
const query_UpdateUserPermissions = ""

const query_DeleteUser = "DELETE FROM main.users WHERE login = :login"

func (e *Data) CreateUser(ctx context.Context, r *misc.DataRequestCreateUser) *misc.DataResponseCreateUser {
	panic("implement me")
}

func (e *Data) GetUserInfo(ctx context.Context, r *misc.DataRequestGetUserInfo) *misc.DataResponseGetUserInfo {
	panic("implement me")
}

func (e *Data) UpdateUserInfo(ctx context.Context, r *misc.DataRequestUpdateUserInfo) *misc.DataResponseUpdateUserInfo {
	panic("implement me")
}

func (e *Data) GetUserList(ctx context.Context, r *misc.DataRequestGetUserList) *misc.DataResponseGetUserList {
	panic("implement me")
}

func (e *Data) GetThreadContent(ctx context.Context, r *misc.DataRequestGetThreadContent) *misc.DataResponseGetThreadContent {
	panic("implement me")
}

func (e *Data) GetPostInfo(ctx context.Context, r *misc.DataRequestGetPostInfo) *misc.DataResponseGetPostInfo {
	panic("implement me")
}

func (e *Data) PostToThread(ctx context.Context, r *misc.DataRequestPostToThread) *misc.DataResponsePostToThread {
	panic("implement me")
}

func (e *Data) GetUserPassHash(ctx context.Context, r *misc.DataRequestGetUserPassHash) *misc.DataResponseGetUserPassHash {
	panic("implement me")
}

func (e *Data) GetUserPermissions(ctx context.Context, r *misc.DataRequestGetUserPermissions) *misc.DataResponseGetUserPermissions {
	panic("implement me")
}

//func (e *DB) CreateUser(c core.DataRequestCreateUser) core.DataResponseCreateUser {
//	tx, err := e.db.BeginTx(context.Background(), &sql.TxOptions{})
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	q, args, err := sqlx.Named(query_CreateUser, c.Request)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	_, err = tx.Exec(q, args...)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	return nil
//}
//func (e *DB) GetUser(c *core.ContainerGetUser) error {
//	q, args, err := sqlx.Named(query_GetUser, c.Request)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	q = e.db.Rebind(q)
//	err = e.db.Get(&c.Response, q, args...)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	return nil
//}
//func (e *DB) GetPassHash(c *core.ContainerGetUserPassHash) error {
//	q, args, err := sqlx.Named(query_GetPassHash, c.Request)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	q = e.db.Rebind(q)
//	err = e.db.Get(&c.Response, q, args...)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	return nil
//}
//func (e *DB) UpdateUser(c *core.ContainerUpdateUser) error {
//	q, args, err := sqlx.Named(query_UpdateUser, c.Request)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	_, err = e.db.Exec(q, args...)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	return nil
//}
//func (e *DB) UpdateUserPassHash(c *core.ContainerUpdateUserPassHash) error {
//	q, args, err := sqlx.Named(query_UpdateUserPassHash, c.Request)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	_, err = e.db.Exec(q, args...)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	return nil
//}
//func (e *DB) DeleteUser(c *core.ContainerDeleteUser) error {
//	q, args, err := sqlx.Named(query_DeleteUser, c.Request)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	_, err = e.db.Exec(q, args...)
//	if err != nil {
//		return errorWrapper(err)
//	}
//
//	return nil
//}
