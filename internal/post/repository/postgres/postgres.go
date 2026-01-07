package post

import "github.com/anthonymartz17/blog_platform_backend.git/internal/entity/post"




type PostPostGres struct{
// db dependecy
}

func New() *PostPostGres{
	return &PostPostGres{}
}

func (r *PostPostGres)Save(post *post.Post) error{
	return nil
}