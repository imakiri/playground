package postgres

import (
	"github.com/imakiri/gorum/types"
	"github.com/jmoiron/sqlx"
	"time"
)

func ThreadCreate(thread types.ModelThread, db *sqlx.DB) error {
	var _, err = db.NamedExec("INSERT INTO app.threads VALUES (:ThreadUUID, :CategoryUUID, :UserUUID, :Name, :DateAdded, :DateLastEdit, :Header)", thread)
	return errWrapper(err)
}
func ThreadGet(threadUUID types.ModelThreadUUID, container *types.ViewThread, db *sqlx.DB) error {
	var err error

	err = db.Get(container, "SELECT category_uuid, name FROM app.threads WHERE thread_uuid = $1", threadUUID)
	if err != nil {
		return errWrapper(err)
	}

	err = db.Get(container.Content.Posts, "SELECT user_uuid, post_uuid, date_added, date_last_edit, content  FROM app.posts WHERE thread_uuid = $1", threadUUID)
	if err != nil {
		return errWrapper(err)
	}

	err = db.Get(container.Content.Users, "SELECT DISTINCT users.user_uuid, registration_date, nick_name, full_name FROM app.posts INNER JOIN app.users ON users.user_uuid = posts.user_uuid WHERE thread_uuid = $1", threadUUID)
	if err != nil {
		return errWrapper(err)
	}

	return errWrapper(err)
}
func ThreadGetList(categoryUUID types.ModelCategoryUUID, container *types.ViewThreadsByCategory, db *sqlx.DB) error {
	var err = db.Get(container, "", categoryUUID)
	return errWrapper(err)
}
func ThreadUpdate(threadUUID types.ModelThreadUUID, container types.ViewThreadUpdate, db *sqlx.DB) error {
	var tx, err = db.Begin()
	if err != nil {
		return errWrapper(err)
	}

	if container.Header != nil {
		_, err = tx.Exec("UPDATE app.threads SET header = $2, date_last_edit = $3 WHERE thread_uuid = $1", threadUUID, container.Header, time.Now().UnixNano())
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errWrapper(e).SetName("Avatar")
			}
			return errWrapper(err).SetName("Avatar")
		}
	}
	if container.Name != nil {
		_, err = tx.Exec("UPDATE app.threads SET name = $2, date_last_edit = $3 WHERE thread_uuid = $1", threadUUID, container.Name, time.Now().UnixNano())
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errWrapper(e).SetName("Fullname")
			}
			return errWrapper(err).SetName("Fullname")
		}
	}
	if container.CategoryUUID != nil {
		_, err = tx.Exec("UPDATE app.threads SET category_uuid = $2, date_last_edit = $3 WHERE thread_uuid = $1", threadUUID, container.CategoryUUID, time.Now().UnixNano())
		if err != nil {
			if e := tx.Rollback(); e != nil {
				return errWrapper(e).SetName("Nickname")
			}
			return errWrapper(err).SetName("Nickname")
		}
	}

	err = tx.Commit()
	return errWrapper(err)
}
func ThreadDelete(threadUUID types.ModelThreadUUID, db *sqlx.DB) error {
	var _, err = db.Exec("DELETE FROM main.app.threads WHERE thread_uuid = $1", threadUUID)
	return errWrapper(err)
}
