package post

import (
	"context"
	"fmt"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
)

//Service uses PostRepository to implement  business logic operations for posts
type Service struct{
	repo PostRepository
}

//New creates a new Service with the given Repository
func New(repo PostRepository) *Service{
	return &Service{repo: repo}

}



//GetPosts retrieves all posts 
func (c *Service)GetPosts(ctx context.Context) ([]entity.Post,error){
	 posts,err:=  c.repo.GetPosts(ctx)

	 if err != nil{
		return nil,fmt.Errorf("Service failed to retrieve posts %w",err)
	 }
return posts,nil
}

//Create creates and saves a new post 
func(c *Service)Create(ctx context.Context,userID, content string)error{
    
	post:= entity.New(userID,content)
  return c.repo.Save(ctx,post)
  
}

