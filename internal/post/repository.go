package post

import (
	"context"
)

//PostRepository  defines the interface for persisting and retrieving posts.
type PostRepository interface{
	GetPosts(ctx context.Context) ([]Post,error)
  Save(ctx context.Context,post *Post)(*Post, error)
}


