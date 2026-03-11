package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/auth"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/middleware"
	entity "github.com/anthonymartz17/blog_platform_backend.git/internal/post"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"

	"github.com/gorilla/mux"
)

//go:generate mockgen -source=http.go -destination=mocks/mock_postcontroller.go -package=mocks

//PostController defines the business logic methods for posts

const (
	msgInvalidBody  = "invalid request body"
	msgInternal     = "internal server error"
	msgEmptyContent = "content cannot be empty"
msgUnauthorized = "unauthorized"
) 
type PostController interface{
	GetPosts(ctx context.Context) ([]entity.Post,error)
	Create(ctx context.Context,userID,content string)error
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
func (h *HTTPHandler)RegisterRoutes(r *mux.Router,authService *auth.Service){
	
	r.HandleFunc("/posts",h.GetPosts).Methods(http.MethodGet)
	
	protected:= r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware(authService))
	
	protected.HandleFunc("/posts",h.Create).Methods(http.MethodPost)


	

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
		log.Printf("failed to decode body %v",err)
		ResponseError(w,http.StatusBadRequest,msgInvalidBody)
		return
	}
	payload.Content= strings.TrimSpace(payload.Content)
	
	
	if payload.Content == ""{
		ResponseError(w,http.StatusBadRequest,msgEmptyContent)
		return
	}
  
	userID,ok:= r.Context().Value(middleware.UserIDKey).(string)
	
	if !ok{
		  log.Println("unable to extract UserIDKey from context")
			 ResponseError(w,http.StatusUnauthorized,msgUnauthorized)
			 return
	}
	


	if err:= h.ctrl.Create(r.Context(),userID,payload.Content); err != nil{

     if errors.Is(err,context.DeadlineExceeded){
			 log.Printf("firebase timeout happened: %v",err)
			 ResponseError(w,http.StatusGatewayTimeout,msgInternal)
		 }

		log.Printf("failed to create post: %v",err)
		ResponseError(w,http.StatusInternalServerError,msgInternal)
		return
	}

	ResponseJSON(w,http.StatusCreated,"Created")
}
