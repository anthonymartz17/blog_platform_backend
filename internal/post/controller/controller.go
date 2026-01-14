package controller

import (
	"context"
	"fmt"

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



//GetPosts retrieves a list of  posts
func (c *Controller)GetPosts(ctx context.Context) error{
	 err:=  c.repo.GetPosts(ctx)

	 if err != nil{
		return fmt.Errorf("Controller failed to retrieve posts %w",err)
	 }
return nil
}

//Create creates and saves a new post 
func(c *Controller)Create(ctx context.Context,post *entity.Post)error{
 return c.repo.Save(ctx,post)
}

