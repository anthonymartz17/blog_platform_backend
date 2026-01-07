package controller

import (
	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
)

type PostRepository interface{
  Save(post *entity.Post) error
}


