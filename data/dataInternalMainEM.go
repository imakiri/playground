package data

import (
	"fmt"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// InternalMain Data Entities
type InternalMainGetUser_1 struct {
	InternalMain
	Request struct {
		InternalMainUserId
		InternalMainUserLogin
	}
	Response struct {
		InternalMainUserAvatar
		InternalMainUserName
	}
}
type InternalMainGetUserPassHash_1 struct {
	InternalMain
	Request struct {
		InternalMainUserLogin
	}
	Response struct {
		InternalMainUserPassHash
	}
}
type InternalMainCreateUser_1 struct {
	InternalMain
	Request struct {
		InternalMainUserLogin
		InternalMainUserAvatar
		InternalMainUserName
		InternalMainUserPassHash
	}
}
type InternalMainDeleteUser_1 struct {
	InternalMain
	Request struct {
		InternalMainUserId
		InternalMainUserLogin
	}
}
type InternalMainUpdateUser_1 struct {
	InternalMain
	Request struct {
		InternalMainUserId
		InternalMainUserLogin
		InternalMainUserName
		InternalMainUserAvatar
	}
}
type InternalMainUpdateUserPassHash_1 struct {
	InternalMain
	Request struct {
		InternalMainUserLogin
		InternalMainUserPassHash
	}
}

// InternalMain Data SQL Methods
func (e *InternalMainDeleteUser_1) SQL() (err error) {
	const query_id = "DELETE FROM main.users WHERE id = ?"
	const query_login = "DELETE FROM main.users WHERE login = ?"
	var query string
	var arg interface{}

	if err = checkRequest(e.Request.Id, e.Request.Login); err != nil {
		return err
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
	return check(err)
}
func (e *InternalMainGetUser_1) SQL() (err error) {
	const query_id = "SELECT name, avatar FROM main.users WHERE id = ?"
	const query_login = "SELECT name, avatar FROM main.users WHERE login = ?"
	var query string
	var arg interface{}

	if err = checkRequest(e.Request.Id, e.Request.Login); err != nil {
		return err
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
	switch e := err.(type) {
	case error:
		return check(e)
	default:
		return
	}
}
func (e *InternalMainGetUserPassHash_1) SQL() (err error) {
	const query = "SELECT passHash FROM main.users WHERE login = ?"

	err = e.SQLX_DB.Get(&e.Response, query, e.Request.Login)
	if err != nil {
		return check(err)
	}

	if e.Response.PassHash == nil {
		return NotFoundError{}
	}

	return
}
func (e *InternalMainCreateUser_1) SQL() (err error) {
	const query = "INSERT INTO main.users (login, name, avatar, passHash) VALUES (?, ?, ?, ?)"

	_, err = e.SQLX_DB.Exec(query, e.Request.Login, e.Request.Name, e.Request.Avatar, e.Request.PassHash)
	return check(err)
}
func (e *InternalMainUpdateUser_1) SQL() (err error) {
	const query_update = "UPDATE main.users"
	const query_set_avatar = "SET avatar = ?"
	const query_set_name = "SET name = ?"
	const query_set_nameAvatar = "SET name = ?, avatar = ?"
	const query_where_id = "WHERE id = ?"
	const query_where_login = "WHERE login = ?"
	var query string
	var args []interface{}

	if err = checkRequest(e.Request.Id, e.Request.Login); err != nil {
		return err
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
	return check(err)
}
func (e *InternalMainUpdateUserPassHash_1) SQL() (err error) {
	const query = "UPDATE main.users SET passHash = ? WHERE login = ?"

	_, err = e.SQLX_DB.Exec(query, e.Request.PassHash, e.Request.Login)
	return check(err)
}
