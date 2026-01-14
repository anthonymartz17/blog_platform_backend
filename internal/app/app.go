package app

import (
	"context"
	"errors"
	"os"

	httpServer "github.com/anthonymartz17/blog_platform_backend.git/internal/http"
	postController "github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	postHandler "github.com/anthonymartz17/blog_platform_backend.git/internal/post/handler"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/post/repository/firebase"
)

//Run initializes dependencies and starts the server
func New() (*httpServer.Server,error){

  ctx:= context.Background()
	fireStoreClient,err:= firebase.NewFirestoreClient(ctx)

	if err != nil{
		return nil,err
	}

	postRepo:= firebase.NewRepo(fireStoreClient)
	postCtrl:= postController.New(postRepo)	
	postHandler:= postHandler.New(postCtrl)	
	

	
	httpRouter:= httpServer.NewRouter()
	postHandler.RegisterRoutes(httpRouter)
  
	port:= os.Getenv("PORT")
	
  if port == ""{
		return nil, errors.New("PORT environment variable not set")
	}

	svr:= httpServer.NewServer(port,httpRouter)

	return svr,nil
   
}


