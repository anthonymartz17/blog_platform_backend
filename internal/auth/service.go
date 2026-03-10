package auth

import (
	"context"

	"firebase.google.com/go/v4/auth"
)


type Service struct{
	client *auth.Client
}
//New creates and return a new AuthService type
func New(c *auth.Client)*Service{
	 return &Service{client:c}
}
//VerifyToken validates idToken and returns retrieved authToken
func (a *Service)VerifyToken(ctx context.Context, idToken string )(*auth.Token,error){
   authToken,err:= a.client.VerifyIDToken(ctx,idToken)
   
	 if err != nil{
		 return nil,err
	 }
	 return authToken,nil
}