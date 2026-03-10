package firebase

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)



func New(ctx context.Context) (*firebase.App,error) {

	serviceAccountPath:= os.Getenv("FIREBASE_CREDENTIALS")

	if serviceAccountPath == "" {
		return nil, fmt.Errorf("FIREBASE_CREDENTIALS environment variable not set")
	}
	

	creds, err := os.ReadFile(serviceAccountPath)

	if err != nil {
		return nil,fmt.Errorf("failed to read firebase credentials: %v", err)
	}

	opt := option.WithAuthCredentialsJSON(option.ServiceAccount,creds)

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return nil,fmt.Errorf("error initializing firebase app: %w", err)
	}

	return app,nil
}
