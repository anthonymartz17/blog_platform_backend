package post

import (
	"time"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/util/uid"
)

type Post struct{
	ID string `json:"id"`
	UserId string `json:"user_id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(userId ,content string) *Post{
 return &Post{
	ID: uid.New(),
	UserId: userId,
	Content: content,
	CreatedAt: time.Now(),
 }
}