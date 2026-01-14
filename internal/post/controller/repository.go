package controller

import (
	"context"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
)

type PostRepository interface{
	GetPosts(ctx context.Context) error
  Save(ctx context.Context,post *entity.Post) error
}


