package models

import "time"

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Type        int       `json:"type"`
	Gender      string    `json:"gender"`
	BirthDay    string    `json:"birthday"`
	PhoneNumber int       `json:"phone_number"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}
