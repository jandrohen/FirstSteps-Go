package models

import (
	"time"
)

// User is the model for the user
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (user *User) AddUser(ID int, Name string, CreatedAt time.Time, status bool) {
	user.ID = ID
	user.Name = Name
	user.CreatedAt = CreatedAt
	user.Status = status
}
