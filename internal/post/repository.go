package post

import (
	"context"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
)

//PostRepository  defines the interface for persisting and retrieving posts.
type PostRepository interface{
	GetPosts(ctx context.Context) ([]entity.Post,error)
  Save(ctx context.Context,post *entity.Post) error
}


