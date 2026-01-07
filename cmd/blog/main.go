package main

import (
	"log"
	"net/http"

	"github.com/anthonymartz17/blog_platform_backend.git/internal/app"
)

/*
Request for testing:

curl -X POST http://localhost:8080/post \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-123",
    "content": "my very first post"
  }'

*/

func main(){
	
	  srv:= app.New()

		log.Printf("✅ Listening on %s", srv.Address())

		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)

	}

}

