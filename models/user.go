package models

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	ID           int64
	Email        string
	Username     string
	PasswordHash string
}

type userContextKey struct{}

func NewUserContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

func UserFromContext(ctx context.Context) *User {
	user, _ := ctx.Value(userContextKey{}).(*User)
	return user
}

func CreateUser(db *sql.DB, user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	query := `INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)`
	result, err := db.Exec(query, user.Email, user.Username, user.PasswordHash)

	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, email, username, password_hash FROM users WHERE email = ?`
	row := db.QueryRow(query, email)

	user := new(User)
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}
