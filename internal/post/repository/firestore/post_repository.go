package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	"google.golang.org/api/iterator"
)
const postsCollection = "posts"
//PostFirebase implements Repository using Firebase as database store
type PostFirebase struct{
client *firestore.Client
}

	//Validates PostFirebase implements Repository interface
	var _ controller.PostRepository = (*PostFirebase)(nil)

//New creats a new PostFirebase
func NewRepo(c *firestore.Client) *PostFirebase{
	return &PostFirebase{
		client: c,
	}
}


//GetPosts retrieves all posts from the "posts" collection in firestore
func (r *PostFirebase)GetPosts(ctx context.Context)([]entity.Post,error){
	iter := r.client.Collection(postsCollection).Documents(ctx)
	defer iter.Stop()

	var posts []entity.Post

for {
	doc, err := iter.Next()
	if err == iterator.Done {
		break
	}

	if err != nil {
		return nil,fmt.Errorf("failed to fetch next Firestore document: %w", err)
	}


	var post entity.Post
	
	if err := doc.DataTo(&post); err != nil {
		return  nil,fmt.Errorf("failed to unmarshal post data (doc ID: %s): %w", doc.Ref.ID, err)
		} 
		
		post.ID = doc.Ref.ID
		posts = append(posts, post)
}
return posts,nil
}


//Save stores a new post to firestore
func (r *PostFirebase)Save(ctx context.Context,post *entity.Post) error{
	//protects against server left hanging waiting
  ctx,cancel:= context.WithTimeout(ctx,5*time.Second)
	defer cancel()
  
	doc:= r.client.Collection(postsCollection).NewDoc()
	post.ID = doc.ID

	_,err:= doc.Set(ctx,post)

	if err != nil {
	return fmt.Errorf("save post to firestore: %w", err)
	}

	return nil
}


func prettyPrintStruct(data any){
	 b, err := json.MarshalIndent(data, "", "  ")
if err != nil {
	log.Println(err)
}

fmt.Println(string(b))
}
