package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	"github.com/gorilla/mux"
)

//go:generate mockgen -source=http.go -destination=mocks/mock_postcontroller.go -package=mocks

//PostController defines the business logic methods for posts
type PostController interface{
	GetPosts(ctx context.Context) ([]entity.Post,error)
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
	r.HandleFunc("/posts",h.GetPosts).Methods(http.MethodGet)
	 r.HandleFunc("/posts",h.Create).Methods(http.MethodPost)

}

//GetPosts retrieves a list of  posts
func (h *HTTPHandler)GetPosts(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()

	posts,err:=  h.ctrl.GetPosts(ctx)

	if err != nil{
		errMsg := fmt.Sprintf("Handler failed to retrieve posts: %v",err)
		ResponseError(w,http.StatusInternalServerError,errMsg)
	  return
	}

	ResponseJSON(w,http.StatusOK,posts)

}

//Create handles http request for creating a post
func (h *HTTPHandler)Create(w http.ResponseWriter, r *http.Request){
	var payload entity.Post
	decoder:= json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err:= decoder.Decode(&payload); err != nil{
		ResponseError(w,http.StatusBadRequest,err.Error())
		return
	}

	payload.Content= strings.TrimSpace(payload.Content)
	
	if payload.Content == ""{
		ResponseError(w,http.StatusBadRequest,"content can not be empty")
		return
	}

	if payload.UserID == ""{
		ResponseError(w,http.StatusBadRequest,"content can not be empty")
		return
	}

	ResponseJSON(w,http.StatusOK,payload)
}