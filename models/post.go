package models

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID         int64
	Title      string
	Content    string
	CategoryID int64
	UserID     int64
	CreatedAt  time.Time
}

func CreatePost(db *sql.DB, post *Post) error {
	if post == nil {
		return errors.New("post is nil")
	}

	query := `INSERT INTO posts (title, content, category_id, user_id) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, post.Title, post.Content, post.CategoryID, post.UserID)

	if err != nil {
		return err
	}

	post.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func GetPostByID(db *sql.DB, postID int64) (*Post, error) {
	query := `SELECT id, title, content, category_id, user_id, created_at FROM posts WHERE id = ?`
	row := db.QueryRow(query, postID)

	post := new(Post)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CategoryID, &post.UserID, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetPostsByCategory(db *sql.DB, categoryID int64) ([]Post, error) {
	rows, err := db.Query("SELECT * FROM posts WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CategoryID, &post.UserID, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetAllPosts(db *sql.DB) ([]Post, error) {
	posts := []Post{}
	rows, err := db.Query("SELECT id, title, content, category_id, user_id, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CategoryID, &post.UserID, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
