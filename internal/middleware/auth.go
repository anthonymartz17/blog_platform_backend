package middleware

import (
	"net/http"
	"github.com/gorilla/mux"
)

// learn about this middleware from gorilla
func AuthMiddleware(verifier AuthVerifier) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
			// 1. Extract Authorization header
			// 2. Validate "Bearer <token>"
			// 3. Verify token using authClient
			// 4. If invalid → write 401 and return
			// 5. If valid → put uid in context
			// 6. Call next.ServeHTTP(w, r.WithContext(newCtx))
       next.ServeHTTP(w,r)
		})
	}
}
// func AuthMiddleware(verifier *auth.Service) mux.MiddlewareFunc {
// 	return func(next http.Handler) http.Handler {
// 			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 					// 1. Read Authorization header
// 					header := r.Header.Get("Authorization")

// 					// 2. Extract Bearer token
// 					tokenString := extractToken(header)

// 					// 3. Verify token
// 					token, err := verifier.VerifyToken(r.Context(), tokenString)
// 					if err != nil {
// 							http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 							return
// 					}

// 					// 4. Store user ID in context
// 					ctx := context.WithValue(r.Context(), "userID", token.UID)

// 					// 5. Continue request
// 					next.ServeHTTP(w, r.WithContext(ctx))
// 			})
// 	}
// }