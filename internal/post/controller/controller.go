package controller

import (
	"context"
	"fmt"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
)

//Controller uses PostRepository to implement  business logic operations for posts
type Controller struct{
	repo PostRepository
}

//New creates a new Controller with the given Repository
func New(repo PostRepository) *Controller{
	return &Controller{repo: repo}

}



//GetPosts retrieves all posts 
func (c *Controller)GetPosts(ctx context.Context) ([]entity.Post,error){
	 posts,err:=  c.repo.GetPosts(ctx)

	 if err != nil{
		return nil,fmt.Errorf("Controller failed to retrieve posts %w",err)
	 }
return posts,nil
}

//Create creates and saves a new post 
func(c *Controller)Create(ctx context.Context,userID, content string)error{
    
	post:= entity.New(userID,content)
  return c.repo.Save(ctx,post)
  
}

