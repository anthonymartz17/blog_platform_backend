package middleware

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

//AuthVerifier defines token validation method
type AuthVerifier interface{
	VerifyToken(ctx context.Context, idToken string )(*auth.Token,error)
}