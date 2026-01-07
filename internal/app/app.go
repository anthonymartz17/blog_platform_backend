package app

import (
	httpServer "github.com/anthonymartz17/blog_platform_backend.git/internal/http"
	postController "github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	postHandler "github.com/anthonymartz17/blog_platform_backend.git/internal/post/handler"
	postRepo "github.com/anthonymartz17/blog_platform_backend.git/internal/post/repository/firebase"
)

//Run initializes dependencies and starts the server
func New() *httpServer.Server{

	port:= ":8080"
  
	postRepo:= postRepo.New()
	postCtrl:= postController.New(postRepo)	
	postHandler:= postHandler.New(postCtrl)	
	

	
	httpRouter:= httpServer.NewRouter()
	postHandler.RegisterRoutes(httpRouter)

	svr:= httpServer.NewServer(port,httpRouter)

	return svr
   
}


