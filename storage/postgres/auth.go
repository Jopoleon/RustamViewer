package postgres

import (
	"time"

	"github.com/pkg/errors"
)

func (db *DB) SetUserSession(user_id int, session string) error {
	_, err := db.DB.Exec("INSERT INTO users_sessions (user_id, session_token, created_at, updated_at) "+
		"VALUES ($1,$2,$3,$4 ) ON CONFLICT (user_id) DO UPDATE SET session_token=excluded.session_token;",
		user_id, session, time.Now(), time.Now())
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) DeleteUserSession(user_id int) error {
	_, err := db.DB.Exec("DELETE FROM users_sessions WHERE user_id=$1",
		user_id)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetUserSession(user_id int) (string, error) {
	var session string
	err := db.DB.Get(&session, "SELECT (session_token) FROM users_session WHERE user_id = $1;", user_id)
	if err != nil {
		return session, errors.WithStack(err)
	}
	return session, nil
}
