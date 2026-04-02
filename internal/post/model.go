package post

import (
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(userID, content string) *Post {
	return &Post{
		UserID:  userID,
		Content: content,
	}
}
