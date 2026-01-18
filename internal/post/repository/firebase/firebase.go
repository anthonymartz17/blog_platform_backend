package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	"google.golang.org/api/iterator"
)

//PostFirebase implements Repository using Firebase as database store
type PostFirebase struct{
client *firestore.Client
}

	//Ensures PostFirebase implements Repository interface
	var _ controller.PostRepository = (*PostFirebase)(nil)

//New creats a new PostFirebase
func NewRepo(c *firestore.Client) *PostFirebase{
	return &PostFirebase{
		client: c,
	}
}


//Save stores a new post to firebase 
func (r *PostFirebase)GetPosts(ctx context.Context)([]entity.Post,error){
	iter := r.client.Collection("posts").Documents(ctx)
	defer iter.Stop()
	var posts []entity.Post

for {
	doc, err := iter.Next()
	if err == iterator.Done {
		break
	}

	if err != nil {
		return nil,fmt.Errorf("Repository error: Failed to iterate: %w", err)
	}


	var post entity.Post
	
	if err := doc.DataTo(&post); err != nil {
		return  nil,fmt.Errorf("Error unmarshaling document data: %v", err)
		} 
		
		post.ID = doc.Ref.ID
		posts = append(posts, post)
}
return posts,nil
}


//Save stores a new post to firebase
func (r *PostFirebase)Save(ctx context.Context,post *entity.Post) error{
	return nil
}

