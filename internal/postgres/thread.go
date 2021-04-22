package postgres

import (
	"time"
)

type ModelCategory struct {
	CategoryUUID string
	Name         string
}

type ModelThread struct {
	ThreadUUID   string
	CategoryUUID string
	UserUUID     string
	Name         string
	DateAdded    int64
	DateLastEdit int64
	Header       string
}

type ViewThreadsByCategory struct {
	ThreadUUID string
	Name       string
}

type ViewUserProfileFromThread struct {
	UserUUID         string
	RegistrationDate int64
	NickName         string
	FullName         string
	Avatar256        []byte
}

type ViewThreadContent struct {
	Users []ViewUserProfileFromThread
	Posts []ViewPostByThreadUUID
}

type ViewPostByThreadUUID struct {
	PostUUID     string
	UserUUID     string
	DateAdded    int64
	DateLastEdit int64
	Content      string
}

type ViewThread struct {
	Category     ModelCategory
	Author       ViewUserProfileFromThread
	Name         string
	DateAdded    int64
	DateLastEdit int64
	Header       string
	Content      ViewThreadContent
}

type ViewThreadUpdate struct {
	CategoryUUID *string
	Name         *string
	Header       *string
}

func ThreadCreate(conn Connection, thread ModelThread) error {
	var _, err = conn.db.NamedExec("INSERT INTO app.threads VALUES (:ThreadUUID, :CategoryUUID, :UserUUID, :Name, :DateAdded, :DateLastEdit, :Header)", thread)
	return err
}
func ThreadGet(conn Connection, threadUUID string, container *ViewThread) error {
	var err error

	err = conn.db.Get(container, "SELECT category_uuid, name FROM app.threads WHERE thread_uuid = $1", threadUUID)
	if err != nil {
		return err
	}

	err = conn.db.Get(container.Content.Posts, "SELECT user_uuid, post_uuid, date_added, date_last_edit, content  FROM app.posts WHERE thread_uuid = $1", threadUUID)
	if err != nil {
		return err
	}

	err = conn.db.Get(container.Content.Users, "SELECT DISTINCT users.user_uuid, registration_date, nick_name, full_name FROM app.posts INNER JOIN app.users ON users.user_uuid = posts.user_uuid WHERE thread_uuid = $1", threadUUID)
	if err != nil {
		return err
	}

	return nil
}
func ThreadGetList(conn Connection, categoryUUID string, container *ViewThreadsByCategory) error {
	return conn.db.Get(container, "", categoryUUID)
}
func ThreadUpdate(conn Connection, threadUUID string, container ViewThreadUpdate) error {
	var tx, err = conn.db.Begin()
	if err != nil {
		return err
	}

	if container.Header != nil {
		_, err = tx.Exec("UPDATE app.threads SET header = $2, date_last_edit = $3 WHERE thread_uuid = $1", threadUUID, container.Header, time.Now().UnixNano())
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return e
			}
			return err
		}
	}
	if container.Name != nil {
		_, err = tx.Exec("UPDATE app.threads SET name = $2, date_last_edit = $3 WHERE thread_uuid = $1", threadUUID, container.Name, time.Now().UnixNano())
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return e
			}
			return err
		}
	}
	if container.CategoryUUID != nil {
		_, err = tx.Exec("UPDATE app.threads SET category_uuid = $2, date_last_edit = $3 WHERE thread_uuid = $1", threadUUID, container.CategoryUUID, time.Now().UnixNano())
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return e
			}
			return err
		}
	}

	err = tx.Commit()
	return err
}
func ThreadDelete(conn Connection, threadUUID string) error {
	var _, err = conn.db.Exec("DELETE FROM main.app.threads WHERE thread_uuid = $1", threadUUID)
	return err
}
