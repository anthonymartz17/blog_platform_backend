package controller

import (
	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
)

//Controller uses a Repository to implement  business logic operations for posts
type Controller struct{
	repo PostRepository
}

//New creates a new Controller with the given Repository
func New(repo PostRepository) *Controller{
	return &Controller{repo: repo}

}

//Create creates and saves a new post 
func(c *Controller)Create(post *entity.Post)error{
 return c.repo.Save(post)
}

