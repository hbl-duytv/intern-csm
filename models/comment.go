package models

import "time"

type Comment struct {
	ID            int       `json:"id"`
	PostID        int       `json:"post_id"`
	CommentatorID int       `json:"commentator_id"`
	Message       string    `json:"message"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
