package middleware

import (
	"context"
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)
type contextKey string
const uidKey contextKey = "uid"


// learn about this middleware from gorilla
func AuthMiddleware(verifier AuthVerifier) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Extract Authorization header
        header:= r.Header.Get("Authorization")
				if !strings.HasPrefix(header, "Bearer ") {
            http.Error(w, "invalid authorization header", http.StatusUnauthorized)
            return
          }

				token:= strings.TrimPrefix(header,"Bearer ")
         
				// 2. Validate "Bearer <token>"
        if token == "" {
           http.Error(w,"missing bearer token",http.StatusUnauthorized)
           return
					}
					
					// 3. Verify token using authClient
					authToken,err:= verifier.VerifyToken(r.Context(),token)
					
					// 4. If invalid → write 401 and return
					if err != nil{
					http.Error(w,"unauthorized",http.StatusUnauthorized)
					return
				}

			// 5. If valid → put uid in contextuidKey
			ctx:= context.WithValue(r.Context(),uidKey,authToken.UID)

        header:= r.Header.Get("Authorization")
				if !strings.HasPrefix(header, "Bearer ") {
            http.Error(w, "invalid authorization header", http.StatusUnauthorized)
            return
          }

				token:= strings.TrimPrefix(header,"Bearer ")
         
				// 2. Validate "Bearer <token>"
        if token == "" {
           http.Error(w,"missing bearer token",http.StatusUnauthorized)
           return
					}
					
					// 3. Verify token using authClient
					authToken,err:= verifier.VerifyToken(r.Context(),token)
					
					// 4. If invalid → write 401 and return
					if err != nil{
					http.Error(w,"unauthorized",http.StatusUnauthorized)
					return
				}

			// 5. If valid → put uid in contextuidKey
			ctx:= context.WithValue(r.Context(),uidKey,authToken.UID)

			// 6. Call next.ServeHTTP(w, r.WithContext(newCtx))
       next.ServeHTTP(w,r.WithContext(ctx))
		})
	}
}

       next.ServeHTTP(w,r.WithContext(ctx))
		})
	}
}
