package firebase

import (
	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
)

//PostFirebase implements Repository using Firebase as database store
type PostFirebase struct{
// db dependecy
}

	//Ensures PostFirebase implements Repository interface
	var _ controller.PostRepository = (*PostFirebase)(nil)

//New creats a new PostFirebase
func New(/*expected dependency*/) *PostFirebase{
	return &PostFirebase{}
}
//Save stores a new post to firebase
func (r *PostFirebase)Save(post *entity.Post) error{
	return nil
}

