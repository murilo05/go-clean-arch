package domain

import "time"

// Change this to your domain
type User struct {
	ID        string
	Document  string
	Name      string
	Email     string
	Age       int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
