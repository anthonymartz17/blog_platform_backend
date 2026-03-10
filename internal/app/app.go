package app

import (
	"context"
	"errors"
	"os"

	httpServer "github.com/anthonymartz17/blog_platform_backend.git/internal/http"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/infrastructure/firebase"
	"github.com/anthonymartz17/blog_platform_backend.git/internal/auth"
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
	
	fireStoreClient,err:= app.Firestore(ctx)
	if err != nil{
		return nil,err
	}

	authClient,err:= app.Auth(ctx)
	if err != nil{
		return nil,err
	}

	authService:= auth.New(authClient)

	postRepo:= postRepository.NewRepo(fireStoreClient)
	postCtrl:= postController.New(postRepo)	
	postHandler:= postHandler.New(postCtrl)	
	
	
	
	httpRouter:= httpServer.NewRouter()
  // httpRouter.Use(auth.Auth(authService))
	postHandler.RegisterRoutes(httpRouter,authService)
  
	port:= os.Getenv("PORT")
	
  if port == ""{
		return nil, errors.New("PORT environment variable not set")
	}

	svr:= httpServer.NewServer(port,httpRouter)

	return svr,nil
   
}


