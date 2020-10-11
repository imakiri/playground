package data

import (
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// Internal_Main Data Entities
type Internal_Main_Method_GetUser_1 struct {
	Internal_Main
	Request struct {
		Internal_Main_User_Id
		Internal_Main_User_Login
	}
	Response []struct {
		Internal_Main_User_Avatar
		Internal_Main_User_Name
	}
}
type Internal_Main_Method_GetUserPassHash_1 struct {
	Internal_Main
	Request struct {
		Internal_Main_User_Login
	}
	Response struct {
		Internal_Main_User_PassHash
	}
}
type Internal_Main_Method_CreateUser_1 struct {
	Internal_Main
	Request struct {
		Internal_Main_User_Login
		Internal_Main_User_Avatar
		Internal_Main_User_Name
		Internal_Main_User_PassHash
	}
}
type Internal_Main_Method_DeleteUser_1 struct {
	Internal_Main
	Request struct {
		Internal_Main_User_Id
		Internal_Main_User_Login
	}
}

// Internal_Main Data ExecuteSQL Methods
func (b *Internal_Main_Method_DeleteUser_1) ExecuteSQL() (err error) {
	const query = "DELETE FROM main.users WHERE id = ? OR login = ?"

	_, err = b.Db.Exec(query, b.Request.Id, b.Request.Login)

	return check(err)
}
func (b *Internal_Main_Method_GetUser_1) ExecuteSQL() (err error) {
	const query = "SELECT name, avatar FROM main.users WHERE login = ? OR  id = ?"

	err = b.Db.Select(&b.Response, query, b.Request.Login, b.Request.Id)
	if err != nil {
		return check(err)
	}

	if len(b.Response) == 0 {
		return Internal_ERROR_NotFound{}
	}

	if len(b.Response) > 1 {
		return Internal_ERROR_IncorrectArgument{}
	}

	return
}
func (b *Internal_Main_Method_GetUserPassHash_1) ExecuteSQL() (err error) {
	const query = "SELECT passHash FROM main.users WHERE login = ?"

	err = b.Db.Select(&b.Response, query, b.Request.Login)
	if err != nil {
		return check(err)
	}

	if b.Response.PassHash == nil {
		return Internal_ERROR_NotFound{}
	}

	return
}
func (b *Internal_Main_Method_CreateUser_1) ExecuteSQL() (err error) {
	const query = "INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)"

	_, err = b.Db.Exec(query, b.Request.Login, b.Request.Name, b.Request.Avatar, b.Request.PassHash)

	return check(err)
}
