package data

import (
	"fmt"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/imakiri/playground/core"
)

// DataInternalMain Data SQL Methods Impl
func requestInternalMainDeleteUser_1(e *core.DataInternalMainDeleteUser_1) {
	const query_id = "DELETE FROM main.users WHERE id = ?"
	const query_login = "DELETE FROM main.users WHERE login = ?"
	var query string
	var arg interface{}
	var err error
	var er core.Err

	if er = checkRequest(e.Request.Id, e.Request.Login); er != nil {
		e.Package.Err = er
		return
	}

	switch {
	case e.Request.Id != 0:
		query = query_id
		arg = e.Request.Id
	case e.Request.Login != "":
		query = query_login
		arg = e.Request.Login
	}

	_, err = e.SQLX_DB.Exec(query, arg)
	e.Package.Err = check(err)
	return
}
func requestInternalMainGetUser_1(e *core.DataInternalMainGetUser_1) {
	const query_id = "SELECT name, avatar FROM main.users WHERE id = ?"
	const query_login = "SELECT name, avatar FROM main.users WHERE login = ?"
	var query string
	var arg interface{}
	var err error
	var er core.Err

	if er = checkRequest(e.Request.Id, e.Request.Login); er != nil {
		e.Package.Err = er
		return
	}

	switch {
	case e.Request.Id != 0:
		query = query_id
		arg = e.Request.Id
	case e.Request.Login != "":
		query = query_login
		arg = e.Request.Login
	}

	err = e.SQLX_DB.Get(&e.Response, query, arg)
	switch err := err.(type) {
	case error:
		e.Package.Err = check(err)
		return
	default:
		return
	}
}
func requestInternalMainGetUserPassHash_1(e *core.DataInternalMainGetUserPassHash_1) {
	const query = "SELECT passHash FROM main.users WHERE login = ?"
	var err error

	err = e.SQLX_DB.Get(&e.Response, query, e.Request.Login)
	if err != nil {
		e.Package.Err = check(err)
		return
	}

	if e.Response.PassHash == nil {
		e.Package.Err = core.NewError(core.DataNotFoundError{}, "")
		return
	}

	return
}
func requestInternalMainCreateUser_1(e *core.DataInternalMainCreateUser_1) {
	const query = "INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)"
	var err error

	_, err = e.SQLX_DB.Exec(query, e.Request.Login, e.Request.Name, e.Request.Avatar, e.Request.PassHash)
	e.Package.Err = check(err)
	return
}
func requestInternalMainUpdateUser_1(e *core.DataInternalMainUpdateUser_1) {
	const query_update = "UPDATE main.users"
	const query_set_avatar = "SET avatar = ?"
	const query_set_name = "SET name = ?"
	const query_set_nameAvatar = "SET name = ?, avatar = ?"
	const query_where_id = "WHERE id = ?"
	const query_where_login = "WHERE login = ?"
	var query string
	var args []interface{}
	var err error
	var er core.Err

	if er = checkRequest(e.Request.Id, e.Request.Login); er != nil {
		e.Package.Err = er
		return
	}

	// Order of args appending is important and depends on query template
	switch {
	case e.Request.Avatar != nil && e.Request.Name == "":
		query = fmt.Sprintf("%s %s", query_update, query_set_avatar)
		args = append(args, e.Request.Avatar)
	case e.Request.Name != "" && e.Request.Avatar == nil:
		query = fmt.Sprintf("%s %s", query_update, query_set_name)
		args = append(args, e.Request.Name)
	case e.Request.Name != "" && e.Request.Avatar != nil:
		query = fmt.Sprintf("%s %s", query_update, query_set_nameAvatar)
		args = append(args, e.Request.Name, e.Request.Avatar)
	}

	switch {
	case e.Request.Id != 0:
		query = fmt.Sprintf("%s %s", query, query_where_id)
		args = append(args, e.Request.Id)
	case e.Request.Login != "":
		query = fmt.Sprintf("%s %s", query, query_where_login)
		args = append(args, e.Request.Login)
	}

	_, err = e.SQLX_DB.Exec(query, args)
	e.Package.Err = check(err)
	return
}
func requestInternalMainUpdateUserPassHash_1(e *core.DataInternalMainUpdateUserPassHash_1) {
	const query = "UPDATE main.users SET passHash = ? WHERE login = ?"
	var err error

	_, err = e.SQLX_DB.Exec(query, e.Request.PassHash, e.Request.Login)
	e.Package.Err = check(err)
	return
}
