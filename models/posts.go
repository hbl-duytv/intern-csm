package models

import "time"

type Posts struct {
	ID          int       `json:"id"`
	Creator     int       `json:"creator"`
	Title       string    `json:"title"`
	Topic       string    `json:"topic"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type TransformPost struct {
	ID          int       `json:"id"`
	Creator     string    `json:"creator"`
	Title       string    `json:"title"`
	Topic       string    `json:"topic"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
