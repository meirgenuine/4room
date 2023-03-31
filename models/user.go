package models

type User struct {
	ID           int64
	Email        string
	Username     string
	PasswordHash string
}

// Implement User-related functions
