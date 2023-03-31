package models

import (
	"database/sql"
	"errors"
	"time"
)

type Comment struct {
	ID        int64
	Content   string
	PostID    int64
	UserID    int64
	CreatedAt time.Time
}

func CreateComment(db *sql.DB, comment *Comment) error {
	if comment == nil {
		return errors.New("comment is nil")
	}

	query := `INSERT INTO comments (content, post_id, user_id) VALUES (?, ?, ?)`
	result, err := db.Exec(query, comment.Content, comment.PostID, comment.UserID)

	if err != nil {
		return err
	}

	comment.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func GetCommentsByPostID(db *sql.DB, postID int64) ([]*Comment, error) {
	query := `SELECT id, content, post_id, user_id, created_at FROM comments WHERE post_id = ?`
	rows, err := db.Query(query, postID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]*Comment, 0)
	for rows.Next() {
		comment := new(Comment)
		err := rows.Scan(&comment.ID, &comment.Content, &comment.PostID, &comment.UserID, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
