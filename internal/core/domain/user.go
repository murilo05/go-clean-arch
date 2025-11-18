package domain

import "time"

// Change this to your domain
type User struct {
	ID        string
	Name      string
	Age       int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
