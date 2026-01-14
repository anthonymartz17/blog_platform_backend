package post

import (
	"context"
	"fmt"
	"net/http"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	"github.com/gorilla/mux"
)

//PostController defines the business logic methods for posts
type PostController interface{
	GetPosts(ctx context.Context) error
	Create(ctx context.Context,post *entity.Post)error
}

//Ensure ctrl.Controller implements the PostController interface.
var _ PostController = (*controller.Controller)(nil)


//HTTPHandler wraps ctrl.PostController and handles post requests
type HTTPHandler struct{
  ctrl PostController
}

//New creates a new HTTPHandler
func New(ctrl PostController)*HTTPHandler{
	return &HTTPHandler{ctrl:ctrl}
}
//RegisterRoutes register post routes
func (h *HTTPHandler)RegisterRoutes(r *mux.Router){
	r.HandleFunc("/post",h.GetPosts).Methods(http.MethodGet)
	 r.HandleFunc("/post",h.Create).Methods(http.MethodPost)

}

//GetPosts retrieves a list of  posts
func (h *HTTPHandler)GetPosts(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()

	err:=  h.ctrl.GetPosts(ctx)
fmt.Println(err)
	// if err != nil{
	//  return fmt.Errorf("Handler failed to retrieve posts %w",err)
	// }
// return nil
}

//Create handles http request for creating posts
func (h *HTTPHandler)Create(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("ok"))
}