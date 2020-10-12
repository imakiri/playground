package data

import (
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// scope (
//     Internal.Main.Method = methods
// ) {
// type methods.GetUser.v1 struct {...}
// func (b *methods.GetUser.V1) Execute.SQL() (err error) {...}
//

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
func (e *Internal_Main_Method_DeleteUser_1) ExecuteSQL() (err error) {
	const query = "DELETE FROM main.users WHERE id = ? OR login = ?"

	_, err = e.SQLX_DB.Exec(query, e.Request.Id, e.Request.Login)

	return check(err)
}
func (e *Internal_Main_Method_GetUser_1) ExecuteSQL() (err error) {
	const query = "SELECT name, avatar FROM main.users WHERE login = ? OR  id = ?"

	err = e.SQLX_DB.Select(&e.Response, query, e.Request.Login, e.Request.Id)
	if err != nil {
		return check(err)
	}

	if len(e.Response) == 0 {
		return Internal_ERROR_NotFound{}
	}

	if len(e.Response) > 1 {
		return Internal_ERROR_IncorrectArgument{}
	}

	return
}
func (e *Internal_Main_Method_GetUserPassHash_1) ExecuteSQL() (err error) {
	const query = "SELECT passHash FROM main.users WHERE login = ?"

	err = e.SQLX_DB.Select(&e.Response, query, e.Request.Login)
	if err != nil {
		return check(err)
	}

	if e.Response.PassHash == nil {
		return Internal_ERROR_NotFound{}
	}

	return
}
func (e *Internal_Main_Method_CreateUser_1) ExecuteSQL() (err error) {
	const query = "INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)"

	_, err = e.SQLX_DB.Exec(query, e.Request.Login, e.Request.Name, e.Request.Avatar, e.Request.PassHash)

	return check(err)
}

//
//}
