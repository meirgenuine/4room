package models

import "time"

type Post struct {
	ID         int64
	Title      string
	Content    string
	CategoryID int64
	UserID     int64
	CreatedAt  time.Time
}

// Implement Post-related functions
