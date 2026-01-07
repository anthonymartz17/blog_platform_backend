package post

import (
	"net/http"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	"github.com/gorilla/mux"
)

//PostController defines the business logic methods for posts
type PostController interface{
	Create(post *entity.Post)error
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
	 r.HandleFunc("/post",h.Create).Methods(http.MethodPost)

}
//Create handles http request for creating posts
func (h *HTTPHandler)Create(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("ok"))
}