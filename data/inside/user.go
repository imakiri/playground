package inside

import (
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// MAIN Data Entities
type MAIN_GetUser_1 struct {
	MAIN
	EXEC
	Request struct {
		MAIN_User_Id
		MAIN_User_Login
	}
	Response []struct {
		MAIN_User_Avatar
		MAIN_User_Name
	}
}
type MAIN_GetUserPassHash_1 struct {
	MAIN
	EXEC
	Request struct {
		MAIN_User_Login
	}
	Response struct {
		MAIN_User_PassHash
	}
}
type MAIN_CreateUser_1 struct {
	MAIN
	EXEC
	Request struct {
		MAIN_User_Login
		MAIN_User_Avatar
		MAIN_User_Name
		MAIN_User_PassHash
	}
}
type MAIN_DeleteUser_1 struct {
	MAIN
	EXEC
	Request struct {
		MAIN_User_Id
		MAIN_User_Login
	}
}

// MAIN Data ExecuteSQL Methods
func (b *MAIN_DeleteUser_1) ExecuteSQL() (err error) {
	const query = "DELETE FROM main.users WHERE id = ? OR login = ?"

	_, err = b.db.Exec(query, b.Request.Id, b.Request.Login)

	return check(err)
}
func (b *MAIN_GetUser_1) ExecuteSQL() (err error) {
	const query = "SELECT name, avatar FROM main.users WHERE login = ? OR  id = ?"

	err = b.db.Select(&b.Response, query, b.Request.Login, b.Request.Id)
	if err != nil {
		return check(err)
	}

	if len(b.Response) == 0 {
		return ERROR_NotFound{}
	}

	if len(b.Response) > 1 {
		return ERROR_IncorrectArgument{}
	}

	return
}
func (b *MAIN_GetUserPassHash_1) ExecuteSQL() (err error) {
	const query = "SELECT passHash FROM main.users WHERE login = ?"

	err = b.db.Select(&b.Response, query, b.Request.Login)
	if err != nil {
		return check(err)
	}

	if b.Response.PassHash == nil {
		return ERROR_NotFound{}
	}

	return
}
func (b *MAIN_CreateUser_1) ExecuteSQL() (err error) {
	const query = "INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)"

	_, err = b.db.Exec(query, b.Request.Login, b.Request.Name, b.Request.Avatar, b.Request.PassHash)

	return check(err)
}
