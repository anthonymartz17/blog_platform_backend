package post

import (
	"context"
	"fmt"
)

// Ensure Service struct implements the PostService interface.
var _ PostService = (*Service)(nil)

//Service uses PostRepository interface to implement  business logic operations for posts
type Service struct{
	repo PostRepository
}

//NewPostService creates a new Service with the given Repository
func NewPostService(repo PostRepository) *Service{
	return &Service{repo: repo}

}





//GetPosts retrieves all posts 
func (c *Service)GetPosts(ctx context.Context) ([]Post,error){
	 posts,err:=  c.repo.GetPosts(ctx)

	 if err != nil{
		return nil,fmt.Errorf("Service failed to retrieve posts %w",err)
	 }
return posts,nil
}

//Create creates and saves a new post 
func(c *Service)Create(ctx context.Context,userID, content string) (*Post, error){
    
	post:= New(userID,content)
  return c.repo.Save(ctx,post)
  
}

