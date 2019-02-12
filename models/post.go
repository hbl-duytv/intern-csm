package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	CreatorID   int       `json:"creator_id"`
	Title       string    `json:"title"`
	Topic       string    `json:"topic"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tag         string    `json:"tag"`
}

type TransformPost struct {
	ID          int       `json:"id"`
	CreatorID   string    `json:"creator_id"`
	Title       string    `json:"title"`
	Topic       string    `json:"topic"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tag         string    `json:"tag"`
}
