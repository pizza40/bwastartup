package user

import "time"

//struct mewakili table users di database
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	Password       string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt time.Time
}