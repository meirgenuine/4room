package models

import "time"

type Comment struct {
	ID        int64
	Content   string
	PostID    int64
	UserID    int64
	CreatedAt time.Time
}
