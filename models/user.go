package models

import "time"

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Type        int       `json:"type"`
	Gender      string    `json:"gender"`
	Birthday    string    `json:"birthday"`
	PhoneNumber int       `json:"phone_number"`
	Status      int       `json:"status"`
	Token       string    `json:"token"`
	Confirm     int       `json:"confirm"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}
