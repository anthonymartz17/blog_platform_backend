package post

import (
	"time"
)

type Post struct{
	ID string `json:"id" firestore:"-"` //for use in struct only since docID is not inside the doc itself.
	UserID string `json:"user_id" firestore:"user_id"`
	Content string `json:"content" firestore:"content"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at"`
}

func New(userID ,content string) *Post{
	now := time.Now()

	return &Post{
		UserID:    userID,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
}