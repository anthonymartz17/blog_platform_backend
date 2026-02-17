package app

import (
	"context"
	"errors"
	"os"

	httpServer "github.com/anthonymartz17/blog_platform_backend.git/internal/http"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/infrastructure/firebase"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/infrastructure/firestore"
	postController "github.com/anthonymartz17/blog_platform_backend.git/internal/post/controller"
	postHandler "github.com/anthonymartz17/blog_platform_backend.git/internal/post/handler"
	postRepository "github.com/anthonymartz17/blog_platform_backend.git/internal/post/repository/firestore"
)

//Run initializes dependencies and starts the server
func New() (*httpServer.Server,error){

  ctx:= context.Background()
	app,err:= firebase.New(ctx)

	if err != nil{
		return nil,err
	}
	fireStoreClient,err:= firestore.NewFirestoreClient(ctx,app)

	if err != nil{
		return nil,err
	}

	

	postRepo:= postRepository.NewRepo(fireStoreClient)
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


