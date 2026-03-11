package main

import (
	"log"
	"net/http"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/app"
	"github.com/joho/godotenv"
)

/*
Request for testing:

curl -X GET http://localhost:8080/post

curl -X POST http://localhost:8080/post \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-123",
    "content": "my very first post"
  }'

*/

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — using system env")
	}
		
	  srv,err:= app.New()

		if err != nil{
			log.Fatal(err)

		}

		log.Printf("✅ Listening on %s", srv.Address())

		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)

	}

}

